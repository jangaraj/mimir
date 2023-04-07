#!/usr/bin/perl
# Copyright 2008 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Modified version of RE2's make_perl_groups.pl.

# Generate table entries giving character ranges
# for POSIX/Perl character classes.  Rather than
# figure out what the definition is, it is easier to ask
# Perl about each letter from 0-128 and write down
# its answer.

@posixclasses = (
	"[:alnum:]",
	"[:alpha:]",
	"[:ascii:]",
	"[:blank:]",
	"[:cntrl:]",
	"[:digit:]",
	"[:graph:]",
	"[:lower:]",
	"[:print:]",
	"[:punct:]",
	"[:space:]",
	"[:upper:]",
	"[:word:]",
	"[:xdigit:]",
);

@perlclasses = (
	"\\d",
	"\\s",
	"\\w",
);

%overrides = (
	# Prior to Perl 5.18, \s did not match vertical tab.
	# RE2 preserves that original behaviour.
	"\\s:11" => 0,
);

sub ComputeClass($) {
  my @ranges;
  my ($class) = @_;
  my $regexp = "[$class]";
  my $start = -1;
  for (my $i=0; $i<=129; $i++) {
    if ($i == 129) { $i = 256; }
    if ($i <= 128 && ($overrides{"$class:$i"} // chr($i) =~ $regexp)) {
      if ($start < 0) {
        $start = $i;
      }
    } else {
      if ($start >= 0) {
        push @ranges, [$start, $i-1];
      }
      $start = -1;
    }
  }
  return @ranges;
}

sub PrintClass($$@) {
  my ($cname, $name, @ranges) = @_;
  print "var code$cname = []rune{  /* $name */\n";
  for (my $i=0; $i<@ranges; $i++) {
    my @a = @{$ranges[$i]};
    printf "\t0x%x, 0x%x,\n", $a[0], $a[1];
  }
  print "}\n\n";
  my $n = @ranges;
  $negname = $name;
  if ($negname =~ /:/) {
    $negname =~ s/:/:^/;
  } else {
    $negname =~ y/a-z/A-Z/;
  }
  return "\t`$name`: {+1, code$cname},\n" .
  	"\t`$negname`: {-1, code$cname},\n";
}

my $gen = 0;

sub PrintClasses($@) {
  my ($cname, @classes) = @_;
  my @entries;
  foreach my $cl (@classes) {
    my @ranges = ComputeClass($cl);
    push @entries, PrintClass(++$gen, $cl, @ranges);
  }
  print "var ${cname}Group = map[string]charGroup{\n";
  foreach my $e (@entries) {
    print $e;
  }
  print "}\n";
  my $count = @entries;
}

print <<EOF;
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// GENERATED BY make_perl_groups.pl; DO NOT EDIT.
// make_perl_groups.pl >perl_groups.go

package syntax

EOF

PrintClasses("perl", @perlclasses);
PrintClasses("posix", @posixclasses);
