# go-gin-demo

[![Actions Status](https://github.com/ducnt114/go-gin-demo/workflows/Go/badge.svg)](https://github.com/ducnt114/go-gin-demo/actions)

API demo with gin and gorm

# Database Schema

```sql
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `hashed_password` varchar(255) NOT NULL,
  `salt` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `token` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` bigint(20) NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `go_gin_demo`.`user` (`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`hashed_password`,`salt`) VALUES ('1','2021-04-23 16:05:50','2021-04-23 16:05:50','','ducnt114','be2b5f4315cfa0167809ce2637a492c9c5411369a4b700dfb7acb7277fc4b9f9018ad093f0d23e43b0002ce8cf64479ed9bc0d4031e13efd96638b73fc8dd04f','randomSalt');
```

# Run

Replace `.env.example` to `.env` and set correct config value in this file.

Then run in terminal:

```
export ENVIRONMENT=LOCAL
go run main.go
```

# API

Login with username and pass

```bash
curl --location --request POST 'http://127.0.0.1:8081/auth/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ducnt114",
    "password": "pass123"
}'
```

Refresh token with value return from login api

```bash
curl --location --request POST 'http://127.0.0.1:8081/auth/v1/refresh-token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "refresh_token": "ksYXHOPtAijYuxhlDnIDtzEzDmHxYifLrxwGyMvOnbicSsRkibsZwceMvVkPnDzitKzlmuCDuPOjEupsevlodoRxVpwYakGHeIPCFunwQjUeTfgbVHwcmYKrsgtanBjS"
}'
```

Access other api with jwt token, ex: get user info api

```bash
curl --location --request GET 'http://127.0.0.1:8081/auth/v1/info' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0MDI4NTAsImlhdCI6MTYxOTM2Njg1MCwidXNlcl9pZCI6MSwidXNlcl9uYW1lIjoiZHVjbnQxMTQifQ.dh9OTYC6WuucnTUSRWnhDXVMMKKP5fa4kbSWgzCXEwxDO4MU1h7CLZKCujhMko6-_POClrfZ077HRgyR9A59yV8xcJcNBQ4K0ZIvh2qp4_rQFGsIeBZWaYNwObRWI2lwO65aTBx82RReGhbVoFKna6k9LG7Aac6_LhCcdqXweR9-ddPDLuXIVziy5W7abbWms5d3P9NKLXsVyEms0hkO7iKmLjfHEeJU8UgmBAIAlU4rg34D9lOFiJHvxRe1V1bRgxbNywujrOR77QwMyWJzIfF9bSmj7Kd9yI-T-zHBFRg85GRiljS5uF3rgiYIoxZT0-8AteX-ejOc8yaS2gLeJGk4j_aWWeFt28Osm4stn4wdn1EvCVMvlO84AsyfKlEPWbqVNAU_VMZvGTvtAbTLGC7E5sP3CzC3obgwBrQPWps-Hpfkl-YlhiE_SidYKm53QIuuwDIYhKTaTJ6O0vka-DEsQ_H2TcJP8vaTZZqC85KQODmApzmuK8CkW7EIr94qn6bJmfyeLRkQG29qcZHpLv3Od9ZylnqXFOGQpHICyjXQmxxAfTileviMmDL54-RWVPdl0Bq2ktp5-25TDaG1xjrhyj5EUid3hhFKUs3rOTNnQIXMolWXsDWrTW79VK7g5YnDy2vp1Szm7p5V5tobHzrFHodZq2ulLepqrWCt94w'
```