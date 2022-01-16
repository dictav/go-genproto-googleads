PKG := github.com/dictav/go-genproto-googleads
GOOGLE_PROTO := google.golang.org/genproto/googleapis/ads/googleads

TARGETS := v7 v8 v9
SRC := googleapis/bazel-bin/google/ads/googleads/$(VERSION)/gapi-ads-googleads-$(VERSION)-go.tar.gz

all: $(TARGETS)

$(TARGETS): googleapis
	-rm -rf build
	make gen VERSION=$@

gen: _build
	-rm -rf $(VERSION) pb/$(VERSION)
	cp -a build/$(PKG)/$(VERSION) $(VERSION)
	cp -a build/$(GOOGLE_PROTO)/$(VERSION) pb/$(VERSION)

_build: $(SRC)
	mkdir build
	tar zxf $(SRC) -C build
	grep -l -R '$(GOOGLE_PROTO)' build/ | \
		xargs sed -i.bak -e 's,$(GOOGLE_PROTO),$(PKG)/pb,'
	find build -name '*.bak' -exec rm {} \;

$(SRC): googleapis
	cd googleapis; bazel build //google/ads/googleads/$(VERSION):gapi-ads-googleads-$(VERSION)-go

googleapis:
	git clone --depth=1 --branch googleads https://github.com/dictav/googleapis

clean:
	-rm -rf build
	-rm -rf googleapis
