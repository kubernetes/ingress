# Custom errors

When the [`custom-http-errors`][cm-custom-http-errors] option is enabled, the Ingress controller configures NGINX so
that it passes several HTTP headers down to its `default-backend` in case of error:

| Header           | Value                                                               |
| ---------------- | ------------------------------------------------------------------- |
| `X-Code`         | HTTP status code returned by the request                            |
| `X-Format`       | Value of the `Accept` header sent by the client                     |
| `X-Original-URI` | URI that caused the error                                           |
| `X-Namespace`    | Namespace where the backend Service is located                      |
| `X-Ingress-Name` | Name of the Ingress where the backend is defined                    |
| `X-Service-Name` | Name of the Service backing the backend                             |
| `X-Service-Port` | Port number of the Service backing the backend                      |
| `X-Request-ID`   | Unique ID that identifies the request - same as for backend service |

A custom error backend can use this information to return the best possible representation of an error page. For
example, if the value of the `Accept` header send by the client was `application/json`, a carefully crafted backend
could decide to return the error payload as a JSON document instead of HTML.

!!! Important
    The custom backend is expected to return the correct HTTP status code instead of `200`.
    NGINX does not change the response from the custom default backend.

An example of such custom backend is available inside the source repository at [images/custom-error-pages][img-custom-error-pages].

See also the [Custom errors][example-custom-errors] example.

[cm-custom-http-errors]: ./nginx-configuration/configmap.md#custom-http-errors
[img-custom-error-pages]: https://github.com/kubernetes/ingress-nginx/tree/main/images/custom-error-pages
[example-custom-errors]: ../../examples/customization/custom-errors
