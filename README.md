# 世界で最も単純なFastCGIサーバ。

FastCGIの挙動がおかしい…PHP-FPMの実装が変なのでは？

といったことを疑った時のために使える、Go言語で書かれた非常に単純なFastCGIサーバ。

```bash
% go run main.go /path/to/socket.sock
2019/09/10 08:48:17 Listen at: /path/to/socket.sock
```

としてLISTENしつつ、例えばnginx側では、

```
#user  nobody;
worker_processes  1;

error_log  logs/error.log debug;
#pid        logs/nginx.pid;


events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
    sendfile        on;
    keepalive_timeout  65;
    server {
        listen       8000;
        server_name  localhost;
        location / {
            root   html;
            index  index.html index.htm;
        }

        # redirect server error pages to the static page /50x.html
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
        location ~ \.php$ {
            root           html;
            fastcgi_pass   unix:/path/to/socket.sock;
            fastcgi_keep_conn on;
            fastcgi_socket_keepalive on;
            fastcgi_index  index.php;
            fastcgi_param  SCRIPT_FILENAME  $document_root/$fastcgi_script_name;
            include        fastcgi_params;
        }
    }
}
```

のような設定ファイルを書いて起動して、 http://localhost:8000/index.php を開いてみることでデバッグができます。
