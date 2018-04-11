# =============================================================================
%define		name	json2csv
%define		version	1.0.0
%define		release	0
%define		branch  master
%define		summary	JSON to CSV Converter
%define		author	John Scherff <jscherff@24hourfit.com>
%define		package	github.com/jscherff/%{name}
%define		gopath	%{_builddir}/go
# =============================================================================

Name:		%{name}
Version:	%{version}
Release:	%{release}%{?dist}
Summary:	%{summary}

License:	ASL 2.0
URL:		https://www.24hourfitness.com
Vendor:		24 Hour Fitness, Inc.
Prefix:		%{_sbindir}
Packager: 	%{packager}
BuildRoot:	%{_tmppath}/%{name}-%{version}-%{release}-root-%(%{__id_u} -n)
Distribution:	el

BuildRequires:    golang >= 1.8.0

%description
The %{name} utility reads a collection of objects in JSON format from a
file or URL and converts those objects to records in a CSV file.

%prep

%build

  export GOPATH=%{gopath}
  export GIT_DIR=%{gopath}/src/%{package}/.git

  go get %{package}
  git checkout %{branch}

  go build -ldflags='-X main.version=%{version}-%{release}' %{package}

%install

  test %{buildroot} != / && rm -rf %{buildroot}/*

  mkdir -p %{buildroot}%{_bindir}
  install -s -m 755 %{_builddir}/%{name} %{buildroot}%{_bindir}/

%clean

  test %{buildroot} != / && rm -rf %{buildroot}/*
  test %{_builddir} != / && rm -rf %{_builddir}/*

%files

  %defattr(-,root,root)
  %{_bindir}/*

%changelog
* Wed Apr 11 2018 - jscherff@24hourfit.com
- Initial build.
