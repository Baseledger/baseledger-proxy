FROM starport/cli

ARG FAUCET_ACC
# override entrypoint from above image
ENTRYPOINT []

WORKDIR /apps

USER root

COPY --chmod=777 . .

ENV FAUCET_ACC=$FAUCET_ACC

EXPOSE 26656 1317 26657

RUN ["starport", "chain", "build"]

# keep the container running
CMD ["tail", "-f", "/dev/null"] 



