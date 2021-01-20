# azl-faas
Serverless function to fetch song lyrics via [AZLyrics]('https://www.azlyrics.com/') using OpenFaaS

### Usage

Deploy:

```bash
$ faas-cli deploy -f azl.yml
```
Invoke:

```text
http://<GATEWAY-IP>:<PORT>/azl?artist=eminem&song=not afraid
```



