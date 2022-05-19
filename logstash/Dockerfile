# https://github.com/elastic/logstash-docker
FROM docker.elastic.co/logstash/logstash:8.2.0
USER root
COPY postgresql-42.3.5.jar /opt/
# Add your logstash plugins setup here
# RUN logstash-plugin install logstash-input-jdbc
# Example: RUN logstash-plugin install logstash-filter-json
RUN logstash-plugin install logstash-filter-aggregate