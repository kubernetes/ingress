# Custom Lua plugins

ingress-nginx uses [https://github.com/openresty/lua-nginx-module](https://github.com/openresty/lua-nginx-module) to run custom Lua code
within Nginx workers. It is recommended to familiarize yourself with that ecosystem before deploying your custom Lua based ingress-nginx plugin.

### Writing a plugin

Every ingress-nginx Lua plugin is expected to have `main.lua` file and all of its dependencies.
`main.lua` is the entry point of the plugin. The plugin manager uses convention over configuration
strategy and automatically runs functions defined in `main.lua` in the corresponding Nginx phase based on their name.

Nginx has different [request processing phases](http://nginx.org/en/docs/dev/development_guide.html#http_phases).
By defining functions with the following names, you can run your custom Lua code in the corresponding Nginx phase:

 - `init_worker`: useful for initializing some data per Nginx worker process
 - `rewrite`: useful for modifying request, changing headers, redirection, dropping request, doing authentication etc
 - `header_filter`: this is called when backend response header is received, it is useful for modifying response headers
 - `body_filter`: this is called when response body is received, it is useful for logging response body 
 - `log`: this is called when request processing is completed and a response is delivered to the client

Check this [`hello_world`](https://github.com/kubernetes/ingress-nginx/tree/main/rootfs/etc/nginx/lua/plugins/hello_world) plugin as a simple example or refer to [OpenID Connect integration](https://github.com/ElvinEfendi/ingress-nginx-openidc/tree/master/rootfs/etc/nginx/lua/plugins/openidc) for more advanced usage.

Do not forget to write tests for your plugin.

### Installing a plugin

There are two options:

  - mount your plugin into `/etc/nginx/lua/plugins/<your plugin name>` in the ingress-nginx pod
  - build your own ingress-nginx image like it is done in the [example](https://github.com/ElvinEfendi/ingress-nginx-openidc/tree/master/rootfs/etc/nginx/lua/plugins/openidc) and install your plugin during image build

Mounting is the quickest option.

### Enabling plugins

Once your plugin is ready you need to use [`plugins` configuration setting](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#plugins) to activate it. Let's say you want to active `hello_world` and `open_idc` plugins, then you set `plugins` setting to `"hello_world, open_idc"`. _Note_ that the plugins will be executed in the given order.
