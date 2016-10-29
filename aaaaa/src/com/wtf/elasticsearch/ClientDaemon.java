package com.wtf.elasticsearch;

import java.net.InetAddress;
import java.net.UnknownHostException;

import org.elasticsearch.action.admin.cluster.health.ClusterHealthResponse;
import org.elasticsearch.action.admin.indices.create.CreateIndexResponse;
import org.elasticsearch.action.admin.indices.settings.get.GetSettingsResponse;
import org.elasticsearch.client.Client;
import org.elasticsearch.client.ClusterAdminClient;
import org.elasticsearch.client.IndicesAdminClient;
import org.elasticsearch.client.transport.TransportClient;
import org.elasticsearch.cluster.health.ClusterHealthStatus;
import org.elasticsearch.cluster.health.ClusterIndexHealth;
import org.elasticsearch.common.settings.Settings;
import org.elasticsearch.common.transport.InetSocketTransportAddress;
import org.elasticsearch.common.unit.TimeValue;
import org.elasticsearch.common.xcontent.XContentBuilder;
import org.elasticsearch.index.engine.Engine.Get;

import com.carrotsearch.hppc.cursors.ObjectObjectCursor;

/**
 * @author wangtengfei-ds
 *
 */
public class ClientDaemon {
	public static Client createClient(String url){
		try{
			Settings settings = Settings.settingsBuilder()
					.put("cluster.name", "logger-es-cluster")
					.put("number_of_replicas", 0)
					.put("number_of_shards", 27)
					.put("transport.tcp.connect_timeout", "90s")
					.build();
			Client client = TransportClient.builder().settings(settings).build().addTransportAddress(new InetSocketTransportAddress(InetAddress.getByName(url),9300));
			return client;
		}catch(Exception e){
			e.printStackTrace();
			return null;
		}
	}
	
	/**
	 * ElasticSearch的API
	 * @param args
	 */
	public static void main(String[] args) {
		Client client = createClient("10.58.65.23");
		clusterAPI(client);
		indicesAPI(client);
	}
	public static void transportClientAPI() throws UnknownHostException{
		// on startup
		TransportClient client1 = TransportClient.builder().build()
		        .addTransportAddress(new InetSocketTransportAddress(InetAddress.getByName("host1"), 9300))
		        .addTransportAddress(new InetSocketTransportAddress(InetAddress.getByName("host2"), 9300));
		
		// on shutdown
		client1.close();
		
		Settings settings2 = Settings.settingsBuilder()
		        .put("cluster.name", "myClusterName").build();
		TransportClient client2 = TransportClient.builder().settings(settings2).build();
		
		Settings settings3 = Settings.settingsBuilder()
		        .put("client.transport.sniff", true).build();
		TransportClient client3 = TransportClient.builder().settings(settings3).build();
		Settings settings4 = Settings.settingsBuilder()
				.put("client.transport.sniff", true)
				.put("client.transport.ping_timeout", TimeValue.timeValueSeconds(5))
				.put("client.transport.nodes_sampler_interval", TimeValue.timeValueSeconds(5))
				.build();
	}
	
	public static void clusterAPI(Client client){
//		开始操作cluster时需要的client
		ClusterAdminClient clusterAdminClient = client.admin().cluster();
//		health API
		ClusterHealthResponse clusterHealthResponse = client.admin().cluster().prepareHealth().get();
		String clusterName = clusterHealthResponse.getClusterName();
		int numberOfDataNode = clusterHealthResponse.getNumberOfDataNodes();
		int numberOfNodes = clusterHealthResponse.getNumberOfNodes();
		System.out.println("=====clusterName====="+clusterName+"=====datanode number====="+numberOfDataNode+"=====numberOfNodes====="+numberOfNodes);
		for (ClusterIndexHealth health : clusterHealthResponse.getIndices().values()) {
			String index = health.getIndex();
			if("cashier-api-2016-10-25".equals(index)){
				int numberOfReplicas = health.getNumberOfReplicas();
				int numberOfShards = health.getNumberOfShards();
				ClusterHealthStatus clusterHealthStatus = health.getStatus();
				String status = clusterHealthStatus.name();
				System.out.println("=====indexName====="+index+"======numberOfShards====="+numberOfShards+"=====numberOfReplicas====="+numberOfReplicas+"=====clusterHealthStatus====="+status);
			}
		}
//		Wait for status
		ClusterHealthResponse healthResponse = client.admin().cluster().prepareHealth().setWaitForYellowStatus().get();
		System.out.println(healthResponse.getStatus().name());
		ClusterHealthResponse healthResponse1 = client.admin().cluster().prepareHealth("cashier-api-2016-10-25").setWaitForGreenStatus().get();
		System.out.println(healthResponse1.getStatus().name());
		ClusterHealthResponse healthResponse2 = client.admin().cluster().prepareHealth("cashier-api-2016-10-25").setWaitForGreenStatus().setTimeout(TimeValue.timeValueSeconds(2)).get();
		System.out.println(healthResponse2.getStatus().name());
	}
	
	public static void indicesAPI(Client client){
//		开始操作index时需要的client
		IndicesAdminClient indicesAdminClient = client.admin().indices();
//		create an index
		CreateIndexResponse createIndexResponse1 = client.admin().indices().prepareCreate("index_name").get();
		System.out.println(createIndexResponse1.getContext().toString());
//		indexSetting
		Settings settings = Settings.builder()
				.put("index.number_of_shards", 5)
				.put("index.number_of_replicas",0)
				.build();
		CreateIndexResponse createIndexResponse2 = client.admin().indices().prepareCreate("index_name").setSettings(settings).get();
		System.out.println("==设置index的setting的执行结果=="+createIndexResponse2.isAcknowledged());
//		put index type mapping
		client.admin().indices().prepareCreate("index_name")   // 	Creates an index called twitter
        .addMapping("tweet", "{\n" +  // It also adds a tweet mapping type.              
                "    \"tweet\": {\n" +
                "      \"properties\": {\n" +
                "        \"message\": {\n" +
                "          \"type\": \"string\"\n" +
                "        }\n" +
                "      }\n" +
                "    }\n" +
                "  }")
        .get();
		CreateIndexResponse createIndexResponse3 = client.admin().indices().prepareCreate("index_name").addMapping("index_type","{\n\"tweet\":{\n\"properties\":{\n\"message\":{\n\"type\": \"string\"\n}\n}\n}\n}").get();
		client.admin().indices().preparePutMapping("twitter")   
		        .setType("user")                                
		        .setSource("{\n" +                              
		                "  \"properties\": {\n" +
		                "    \"name\": {\n" +
		                "      \"type\": \"string\"\n" +
		                "    }\n" +
		                "  }\n" +
		                "}")
		        .get();
		// You can also provide the type in the source document
		client.admin().indices().preparePutMapping("twitter")
		        .setType("user")
		        .setSource("{\n" +
		                "    \"user\":{\n" +                        
		                "        \"properties\": {\n" +
		                "            \"name\": {\n" +
		                "                \"type\": \"string\"\n" +
		                "            }\n" +
		                "        }\n" +
		                "    }\n" +
		                "}")
		        .get();
		client.admin().indices().preparePutMapping("twitter")   
        .setType("tweet")                               
        .setSource("{\n" +                              
                "  \"properties\": {\n" +
                "    \"user_name\": {\n" +
                "      \"type\": \"string\"\n" +
                "    }\n" +
                "  }\n" +
                "}")
        .get();
//		refresh index
		client.admin().indices().prepareRefresh().get(); 
		client.admin().indices()
		        .prepareRefresh("twitter")               
		        .get();
		client.admin().indices()
		        .prepareRefresh("twitter", "company")   
		        .get();
//		get index setting
		GetSettingsResponse response = client.admin().indices()
		        .prepareGetSettings("index_name", "index_type").get();                           
		for (ObjectObjectCursor<String, Settings> cursor : response.getIndexToSettings()) { 
		    String index = cursor.key;                                                      
		    Settings settings3 = cursor.value;                                               
		    Integer shards = settings3.getAsInt("index.number_of_shards", null);             
		    Integer replicas = settings3.getAsInt("index.number_of_replicas", null);
		    System.out.println("==index=="+index+"==shards=="+shards+"==replicas=="+replicas);
		}
//		update index setting
		client.admin().indices().prepareUpdateSettings("twitter")   
		        .setSettings(Settings.builder()                     
		                .put("index.number_of_replicas", 0)
		        )
		        .get();
		
	}
	
}
