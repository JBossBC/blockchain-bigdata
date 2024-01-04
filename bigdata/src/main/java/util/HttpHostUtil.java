package util;

import org.apache.http.HttpHost;

import java.util.ArrayList;
import java.util.List;

public class HttpHostUtil {

    public static HttpHost[] httpHosts(String urls) {
        List<HttpHost> httpHosts = new ArrayList<>();
        for (String url : urls.split("\\s*,\\s*")) {
            String[] split = url.trim().split(":");
            String http = split[0];
            String ip = split[1].substring(2);
            int port = Integer.parseInt(split[2]);
            httpHosts.add(new HttpHost(ip, port, http));
        }
        HttpHost[] hosts =new HttpHost[1];
        hosts=httpHosts.toArray(hosts);
        return hosts;
    }
}
