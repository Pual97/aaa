func TestMatchIPV4(t *testing.T) {
	t.Parallel()
	server := setupTest(awsIPResponse{
		Prefixes: []prefixEntry{
			{IPV4Prefix: "192.168.0.0/24"},
		},
	})
	defer server.Close()

	ctx := context.Background()
	ips, _ := newAWSIPs(ctx, serverIPRanges(server), time.Hour, nil)
	err := ips.tryUpdate(ctx)
	assertEqual(t, err, nil)
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.0")))
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.1")))
	assertEqual(t, false, ips.contains(net.ParseIP("192.169.0.0")))
}


func TestMatchIPV4(t *testing.T) {
	t.Parallel()
	server := setupTest(awsIPResponse{
		Prefixes: []prefixEntry{
			{IPV4Prefix: "192.168.0.0/24"},
		},
	})
	defer server.Close()

	ctx := context.Background()
	ips, _ := newAWSIPs(ctx, serverIPRanges(server), time.Hour, nil)
	err := ips.tryUpdate(ctx)
	assertEqual(t, err, nil)
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.0")))
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.1")))
	assertEqual(t, false, ips.contains(net.ParseIP("192.169.0.0")))
}

func TestMatchIPV4(t *testing.T) {
	t.Parallel()
	server := setupTest(awsIPResponse{
		Prefixes: []prefixEntry{
			{IPV4Prefix: "192.168.0.0/24"},
		},
	})
	defer server.Close()

	ctx := context.Background()
	ips, _ := newAWSIPs(ctx, serverIPRanges(server), time.Hour, nil)
	err := ips.tryUpdate(ctx)
	assertEqual(t, err, nil)
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.0")))
	assertEqual(t, true, ips.contains(net.ParseIP("192.168.0.1")))
	assertEqual(t, false, ips.contains(net.ParseIP("192.169.0.0")))
}
