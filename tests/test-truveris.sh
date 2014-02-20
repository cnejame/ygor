#!/bin/sh

. ./_functions.sh

# $1 command
cmd() {
	sleep 0.1
	echo "$@"
	sleep 0.1
}

# $1 command
test_line() {
	cmd "$@" \
		| ../ygor-truveris/ygor-truveris --nickname=whygore \
		2> test.stderr \
		> test.output
	if [ "$?" != 0 ]; then
		fail "wrong return code (check test.stderr)"
	fi
}


announce "auto-joins"
test_line ""
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
EOF
assert_output && pass


announce "unknown chatter"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :blabla"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
EOF
assert_output && pass


announce "set a new alias"
rm -f aliases.cfg
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: alias blabla play stuff.ogg"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
PRIVMSG #test :ok (created)
EOF
assert_output && pass


announce "get this new alias"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: alias blabla"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
PRIVMSG #test :'blabla' is an alias for 'play stuff.ogg'
EOF
assert_output && pass


announce "get unknown alias"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: alias whatevs"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
PRIVMSG #test :error: unknown alias
EOF
assert_output && pass


announce "list all known aliases"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: aliases"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
PRIVMSG #test :known aliases: blabla
EOF
assert_output && pass


announce "test alias with percent sign"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: alias 60% play stuff"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: alias 60%"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
PRIVMSG #test :'60%' is an alias for 'play stuff'
EOF
assert_output && pass


announce "test say stuff"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: say stuff"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
[SQS-SendToMinion] say stuff
EOF
assert_output && pass


announce "test use alias"
test_line ":jimmy!dev@truveris.com PRIVMSG #test :whygore: 60%"
cat > test.expected <<EOF
JOIN #test
JOIN #ygor
[SQS-SendToMinion] play stuff
EOF
assert_output && pass


cleanup