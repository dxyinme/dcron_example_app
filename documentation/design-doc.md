# Design Doc

## Background

This application just a example application for [Dcron](https://github.com/libi/dcron). The work this app want to handle, which is managing your cron task to execute SQLs on many databases. You can register your databases, your SQL tasks.

## REST API

This project use swaggo over gin framework to provide an easy REST API.

[swagger docs](../app/docs/)

## Design 

Use Dcron to managed cron tasks. And use inner-call to boardcast the cron-task CRUD in cluster. Use MySQL to save databases and tasks data. Use redis pub-sub to implement inner call.

## Inner Call

Every instance will subscribe to the same channel in redis. Any one instance receive a cron-task create / update / remove , this instance will publish the request to this channel.
