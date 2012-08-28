TEST_ARTIFACTS = \
  test-compilation-stamp \
  test-configuration-stamp \
  test-installation-stamp \
  test-merge-stamp \
  test-stamp \
  thrift \
  thrift-fetch-stamp \
  tree \
  tree-stamp

export THRIFT=$(CURDIR)/thrift
export THRIFT4GO=$(CURDIR)/../..

all: test

test: test-stamp

test-stamp: test-installation-stamp
	touch $@

test-merge-stamp: thrift-fetch-stamp
	cd $(THRIFT4GO) && bash ./scripts/merge_and_build.sh -b
	touch $@

test-configuration-stamp: test-merge-stamp
	cd $(THRIFT) && ./configure \
		--with-go \
		--without-c_glib \
		--without-erlang \
		--without-haskell \
		--without-java \
		--without-perl \
		--without-php \
		--without-python \
		--without-ruby
	touch $@

test-compilation-stamp: test-configuration-stamp
	cd $(THRIFT) && make
	touch $@

test-installation-stamp: test-compilation-stamp tree-stamp
	cd $(THRIFT) && DESTDIR=$(CURDIR)/tree make install
	touch $@

tree-stamp:
	mkdir -vp tree
	touch $@

thrift-fetch-stamp:
	svn co http://svn.apache.org/repos/asf/thrift/trunk thrift
	touch $@

clean:
	rm -rf $(TEST_ARTIFACTS)

.PHONY: test
