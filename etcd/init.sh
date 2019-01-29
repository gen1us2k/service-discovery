curl -XPUT http://127.0.0.1:23791/v2/keys/skydns/local/skydns/east/production/rails/1 \
	    -d value='{"host":"service1.example.com","port":8080}'
curl -XPUT http://127.0.0.1:23791/v2/keys/skydns/local/skydns/west/production/rails/2 \
	    -d value='{"host":"service2.example.com","port":8080}'
curl -XPUT http://127.0.0.1:23791/v2/keys/skydns/local/skydns/east/staging/rails/4 \
	    -d value='{"host":"10.0.1.125","port":8080}'
curl -XPUT http://127.0.0.1:23791/v2/keys/skydns/local/skydns/east/staging/rails/6 \
	    -d value='{"host":"2003::8:1","port":8080}'
curl -XPUT http://127.0.0.1:23791/v2/keys/skydns/config -d value='{"dns_addr":"127.0.0.1:5354","ttl":3600, "nameservers": ["8.8.8.8:53","8.8.4.4:53"]}'



