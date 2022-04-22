FROM golang:1.18 

# Create and change the building directory to /build/
WORKDIR /build

# Add necessary components to build executable
ADD go.mod go.sum Makefile ./ 
ADD ./internal ./internal
ADD ./vendor ./vendor
ADD ./cmd ./cmd

# Run the Makefile to create the executable in /build/bin/rest-api-template
RUN make build

# Copy and run the executable in /usr/local/bin/rest-api-template 
RUN cp /build/bin/rest-api-template /usr/local/bin/rest-api-template
CMD ["/usr/local/bin/rest-api-template"]
