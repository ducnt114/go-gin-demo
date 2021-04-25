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

if output like this then run api ok:

```bash
$ go run main.go                                                                    
{"level":"info","ts":1619370861.354134,"caller":"conf/config.go:45","msg":"Environment: LOCAL"}
{"level":"info","ts":1619370861.3556976,"caller":"go-gin-demo/main.go:20","msg":"Starting api at: :8081"}
```

# API

## Login with username and pass

```bash
curl --location --request POST 'http://127.0.0.1:8081/auth/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ducnt114",
    "password": "pass123"
}'
```

Response:

```json
{
    "meta": {
        "code": 200,
        "message": "success"
    },
    "data": {
        "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0MDY4NjYsImlhdCI6MTYxOTM3MDg2NiwidXNlcl9pZCI6MSwidXNlcl9uYW1lIjoiZHVjbnQxMTQifQ.IVUsNCNe-ShNN7DAmdsFHwzNfsf71o5VfnXsw0T76Kls2yRVtNYWGfDAj9ScwTz7_FLUD0-N4Q7x4BxNBgUL1vDkAx7y5sA4A7GcSu5CbbR1cuHe-ITthI-AQ3RhVChA-hb89jrojxByrva40Ky7x3HNRiYUDjGzyTOMi0wGNFnTTvhsRE7tezaAwHyCWqdMXq4LfVyKoqKaXstB8FEvQecn9KHob3Y9AxzumaLUbDSniCvj40estirf5mK9ydNtiDyAYO44_TGTVgJ_Zqn9S2ycc-9yuSkpw0TDOfO1w1XjFklIsM6mR-_j5BnWflAJs4DOcEWvxXWsR_WawAAd8z6XGD2-klb8Eb_9HxS1VYb1SeOS9bq8g540BCKn3eKsJBHpwROnaJuW4sRVAWzpb6xevhzIzpm1Op4rqGIH1HFKt4M1jQIfAvCSzG-IyZUcQumDsNvQw6ZHi6aA-z4gmX9W2__fBKVIs_oKqey6sBGmwi8F0wiGZCqTrV_lVpjtNt0pLB_tirfDzZ2IXIj7DOwnewCmI3dvPIGsq8w4eKS91lRTiVfdd6CyHaRgstnKpaZtl1aSlk4K3Xds988Sqt90W2H_pJPiIn59RuzUX-L16VjDiIs9zHkABHGrDq4WvlN6MFfzZn0HIN6zYfZiWJQ4hYTywYEFrx3uGicBO3w",
        "refresh_token": "ksYXHOPtAijYuxhlDnIDtzEzDmHxYifLrxwGyMvOnbicSsRkibsZwceMvVkPnDzitKzlmuCDuPOjEupsevlodoRxVpwYakGHeIPCFunwQjUeTfgbVHwcmYKrsgtanBjS"
    }
}
```

## Refresh token with value return from login api

```bash
curl --location --request POST 'http://127.0.0.1:8081/auth/v1/refresh-token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "refresh_token": "ksYXHOPtAijYuxhlDnIDtzEzDmHxYifLrxwGyMvOnbicSsRkibsZwceMvVkPnDzitKzlmuCDuPOjEupsevlodoRxVpwYakGHeIPCFunwQjUeTfgbVHwcmYKrsgtanBjS"
}'
```

Response 

```json
{
    "meta": {
        "code": 200,
        "message": "success"
    },
    "data": {
        "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0MDY4OTUsImlhdCI6MTYxOTM3MDg5NSwidXNlcl9pZCI6MSwidXNlcl9uYW1lIjoiZHVjbnQxMTQifQ.KpGYhpUJhC2ZuK0CbHyXypuMQC7-9iihXpGJ4_ME8UTgz1ofOZlziUHbU7ZlG3IIz5ELFdkoNDA-wUBQQMnhJcwjbo-pTBo2B2hpVe64w7yvavCoK-ZH0A3UR39lIqlEZl844unFkt6nlNdCkSdom3D6XXIp1KX-qcBDXITtNlJoIOIYODbJQrzVZ6vUFZHBG5_8GiEkwhulOYH5UjV0PV004E8jCFqLls-tVGZwZ8yN2Jmn5JQEtOblg_mYTV8E4OM-whgW-d54keLY70q88FdfEim9gVycO7JN0u_Nu3Z3YVzgb_yoi6LNYdQbwllS0DwjHzDiG09JBXXEJG7nu8BVhuCkUp1tuF95VCqlGkWCRxJkZnCWwygBenCbKP6Qgn7y9TEOMhI1yj1A4u8ipEO5BDEckzKPlo7ihPHyubs84xvGJJwSQ-4qJYKzJUStOSu-B7PPEz34ZgaMI4laJnIRqjxV1rUWXoQmbi5MaYqDVWTlWEgwJj57sKtgdup9UhZCOD3WJhbVmPKimZoWtlP27vbjKPzlxq9xXZHjkICpDtlHBWkMTG-i9ESRQBx6Ee6UCpnm_lMm0G7lfTp7UKsiUFrKpMS6sV2m-BLV69_QAyO3BLUYciE__Pgq_HaWbi02b6eWZ-w7PEAtKfjl8jqF_6qIrg7wqFCLVYLyEHg",
        "refresh_token": "qTqMDEdrweXxRYjLdGNTdgdyBenaTxBMfhyjJEkyYhlOyjkrjlJZbmTDsNGcUKixRVKbtUlMtznpVrwksNCgDtgogFsTZNWBZAMawTlWKFVrARBzvoAkeGOllXrgyolk"
    }
}
```

## Access other api with jwt token, ex: get user info api

```bash
curl --location --request GET 'http://127.0.0.1:8081/auth/v1/info' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk0MDI4NTAsImlhdCI6MTYxOTM2Njg1MCwidXNlcl9pZCI6MSwidXNlcl9uYW1lIjoiZHVjbnQxMTQifQ.dh9OTYC6WuucnTUSRWnhDXVMMKKP5fa4kbSWgzCXEwxDO4MU1h7CLZKCujhMko6-_POClrfZ077HRgyR9A59yV8xcJcNBQ4K0ZIvh2qp4_rQFGsIeBZWaYNwObRWI2lwO65aTBx82RReGhbVoFKna6k9LG7Aac6_LhCcdqXweR9-ddPDLuXIVziy5W7abbWms5d3P9NKLXsVyEms0hkO7iKmLjfHEeJU8UgmBAIAlU4rg34D9lOFiJHvxRe1V1bRgxbNywujrOR77QwMyWJzIfF9bSmj7Kd9yI-T-zHBFRg85GRiljS5uF3rgiYIoxZT0-8AteX-ejOc8yaS2gLeJGk4j_aWWeFt28Osm4stn4wdn1EvCVMvlO84AsyfKlEPWbqVNAU_VMZvGTvtAbTLGC7E5sP3CzC3obgwBrQPWps-Hpfkl-YlhiE_SidYKm53QIuuwDIYhKTaTJ6O0vka-DEsQ_H2TcJP8vaTZZqC85KQODmApzmuK8CkW7EIr94qn6bJmfyeLRkQG29qcZHpLv3Od9ZylnqXFOGQpHICyjXQmxxAfTileviMmDL54-RWVPdl0Bq2ktp5-25TDaG1xjrhyj5EUid3hhFKUs3rOTNnQIXMolWXsDWrTW79VK7g5YnDy2vp1Szm7p5V5tobHzrFHodZq2ulLepqrWCt94w'
```

Response:

```json
{
    "meta": {
        "code": 200,
        "message": "success"
    },
    "data": {
        "user_id": 1,
        "user_name": "ducnt114"
    }
}
```