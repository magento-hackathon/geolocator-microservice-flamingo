FROM aoepeople/scratch-go-env

ADD config /config
ADD bin/geolocatorservice_unix bin/geolocatorservice_unix

ENTRYPOINT ["bin/geolocatorservice_unix"]

EXPOSE 3322

CMD ["serve"]
