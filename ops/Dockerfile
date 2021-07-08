FROM starport/cli

# override entrypoint from above image
ENTRYPOINT []

WORKDIR /apps

USER root

ARG DB_HOST
ARG DB_UB_USER
ARG DB_UB_PWD

COPY --chmod=777 . .

EXPOSE 26656 1317

ENV DB_HOST=$DB_HOST
ENV DB_UB_USER=$DB_UB_USER
ENV DB_UB_PWD=$DB_UB_PWD
ENV DB_UB_NAME=ub
ENV DB_SSLMODE=disable
ENV DB_BASELEDGER_NAME=baseledger
ENV DB_BASELEDGER_USER=baseledger
ENV NATS_URL=nats:4222

# CMD ["build"]
RUN ["starport", "build"]

# keep the container running
CMD ["tail", "-f", "/dev/null"] 



