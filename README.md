# GoneyPot

<p align="center">
<img width="400" src="https://user-images.githubusercontent.com/6976628/85920077-e5d4d500-b870-11ea-8ee7-b51d905e0032.png" />
</p>

A golang CLI to setup [honey pots](https://en.wikipedia.org/wiki/Honeypot_(computing)). For now, GoneyPot can only be used for setting up passive honey pots, which can't be interacted with and just listen to scans on the network.

## Features

* Listens for scanning on all TCP and UDP ports
* Detects ICMP & ping messages

## Todo

* Grafana Dashboard with Prometheus Metrics
    * Email alerting
* Signature-based intrusion detection
    * Severity levels
* Service emulation (interactive honey pot)
