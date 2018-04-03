# geo-trouble

This golang server is loading informations from csv to mongodb and exploit the datas to return the number & color of the accidents in a given French city / date(from -> to)

### Rules for colors : 

The coloration rules are set into the ```config.toml```.
You can modify the range of it.

### Prerequisites : 

Install Mongodb (or dockerise it)

Use dep for packages (https://github.com/golang/dep)
```
dep ensure install
```

### Test it :

You can modify credentials into the ```config.toml``` for ports and conenctions.

```
Go run server.go
```

### Utility Links :

```
http://localhost:9090/city/Paris/dates/2018-11-20/2018-11-29 -> 13 accidents
http://localhost:9090/city/Paris/dates/2018-11-20/2018-11-29/color -> blue
http://localhost:9090/city/Paris/dates/2018-11-10/2018-11-29/color -> green
http://localhost:9090/city/Montpellier/dates/2018-09-20/2018-11-29/color -> orange
http://localhost:9090/city/Montpellier/dates/2018-02-20/2018-11-29/color -> red
```

Dao, models, config
Fonctional test are present

### Author :
Laurent Loukopoulos