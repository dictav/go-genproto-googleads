PKG := github.com/dictav/go-genproto-googleads
GOOGLE_PROTO := google.golang.org/genproto/googleapis/ads/googleads

TARGETS := v13 v14 v15
SRC := googleapis/bazel-bin/google/ads/googleads/$(VERSION)/gapi-ads-googleads-$(VERSION)-go.tar.gz

all: $(TARGETS)

$(TARGETS): googleapis
	-rm -rf build
	$(MAKE) gen VERSION=$@

gen: build
	-rm -rf $(VERSION) pb/$(VERSION)
	cp -a build/$(PKG)/$(VERSION) $(VERSION)
	cp -a build/$(GOOGLE_PROTO)/$(VERSION) pb/$(VERSION)

build: $(SRC)
	mkdir tmp
	tar zxf $(SRC) -C tmp
	grep -l -R '$(GOOGLE_PROTO)' tmp/ | \
		xargs sed -i -e 's,$(GOOGLE_PROTO),$(PKG)/pb,'
	mv tmp/$(GOOGLE_PROTO)/$(VERSION)/resources/experiment_arm.pb.go \
		tmp/$(GOOGLE_PROTO)/$(VERSION)/resources/experimentarm.pb.go
	mv tmp build

$(SRC): googleapis
	cd googleapis; git reset --hard
	./googleads-go.patch.sh $(PKG) $(VERSION)
	cd googleapis; bazel build --cxxopt=-std=c++14 --host_cxxopt=-std=c++14 //google/ads/googleads/$(VERSION):gapi-ads-googleads-$(VERSION)-go

googleapis:
	git clone --depth=1 https://github.com/googleapis/googleapis

clean:
	-rm -rf tmp
	-rm -rf build
	-rm -rf googleapis
