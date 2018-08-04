FROM flemay/golang:1-stretch AS builder
COPY . /go/src/{{cookiecutter.app_go_package_path}}/
WORKDIR /go/src/{{cookiecutter.app_go_package_path}}
ENV IS_SCRATCH_IMAGE true
RUN make _deps _test _build

FROM scratch
LABEL maintainer "@{{cookiecutter.git_username}}"
COPY --from=builder /go/src/{{cookiecutter.app_go_package_path}}/bin/{{cookiecutter.git_repo_name}} /
ENTRYPOINT [ "/{{cookiecutter.git_repo_name}}" ]
