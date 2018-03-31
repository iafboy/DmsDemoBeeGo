FROM golang

#author
MAINTAINER Dabing

#add beego and other package
ADD github.com /go/src/github.com

#build bee tool script
ADD build.sh /build.sh
RUN chmod +x /build.sh
RUN /build.sh

# install beego
RUN go get github.com/astaxie/beego

#add bee tool to PATH
ENV PATH $PATH:$GOPATH/bin

#copy project
ADD FAWCarDMS /go/src/FAWCarDMS


#start project script
ADD run.sh /
RUN chmod +x /run.sh

EXPOSE 8080

CMD ["/run.sh"]