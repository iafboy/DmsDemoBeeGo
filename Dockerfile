FROM golang

#author
MAINTAINER Dabing

#add beego and other package
ADD github.com /go/src/github.com

#build bee tool script
ADD build.sh /build.sh
RUN chmod +x /build.sh
RUN /Build.sh

# install beego
RUN go get github.com/astaxie/beego

#add bee tool to PATH
ENV PATH $PATH:$GOPATH/bin

#generate project
ADD autoCreatePrj.sh /autoCreatePrj.sh
RUN chmod +x /autoCreatePrj.sh
RUN /autoCreatePrj.sh


#start project script
ADD run.sh /

RUN chmod +x /run.sh

EXPOSE 8080

CMD ["/run.sh"]