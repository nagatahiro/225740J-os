#!/usr/bin/perl

use strict;
use warnings;
use Socket;
use Sys::Hostname;

my $host = hostname;
my $iface;
my %interfaces = ();

foreach my $line (`ifconfig`) {
    if ($line =~ /^([a-zA-Z0-9]+):/) {
        $iface = $1;
    } elsif ($line =~ /inet (\d+\.\d+\.\d+\.\d+)/) {
        push @{$interfaces{$iface}}, $1;
    }
}

foreach $iface (keys %interfaces) {
    print "The name of this network interface is $iface.\n";
    print "  All the IP addresses of this network interface are: \n";
    foreach my $addr (@{$interfaces{$iface}}) {
        print "    $addr\n";
    }
}