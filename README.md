# azl-faas
Serverless function to fetch song lyrics via [AZLyrics]('https://www.azlyrics.com/') using OpenFaaS

For this function, I am making use of [golang-http-template](https://github.com/openfaas/golang-http-template). Specifically the `golang-middleware`

```bash
$ faas-cli template store pull golang-middleware
```
### Usage

Deploy:

```bash
$ faas-cli deploy -f azl.yml
```
Invoke:

```text
http://<GATEWAY-IP>:<PORT>/azl?artist=eminem&song=not afraid
```
Output:

![azl-faas](https://github.com/prashantkhandelwal/azl-faas/blob/master/screenshots/azl-screenshot.png)