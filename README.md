# Geolocation Microservice in Flamingo (www.flamingo.me)

Simple Geolocationing Microservice in Flamingo with a base set of providers (prone to changes, currently www.ipstack.com)

## Getting Started

- Install Go
- Obtain API Key from www.ipstack.com and edit the config/config.yml accordingly
- fire up ```go run main.go serve```
- Its running, should be reachable at ```http://localhost:3322/geolocation/[insert IP Address here]```

### Prerequisites

You will need Golang in version 1.12.1 installed.

### Configuration

The config/config.yml contains a provider configuration with api key and url (provider specific), which could grow by
supported providers in the future. New Providers have to implement the LocationProvider interface to work.

## Running the tests

Use the following to run the tests:

```
go test ./... -v
```

## Built With

* [Flamingo](https://go.aoe.com/#Home) - Scalable frontend framework for your headless microservice architecture

## Authors

* **Jens Sadowski** - *Initial work* - [resubaka](https://github.com/resubaka)
* **Julien Wiedner** - *Initial work* - [juce93](https://github.com/juce93)
* **Joachim Adomeit** - *Initial work* - [jadhub](https://github.com/jadhub)

## License

This project is licensed under the Open Software License v. 3.0 (OSL-3.0) - see the [LICENSE](LICENSE) file for details

