package crypt

//cFcrypt
//
//This is the go-implementation of
//
//https://github.com/ptt/pttbbs/blob/master/common/sys/crypt.c
//commit: 6bdd36898bde207683a441cdffe2981e95de5b20
//
//Params
//
//	buf: input-buffer (raw)
//	salt: salt
//	buff: output-buffer (encrypted)
func cFcrypt(buf []uint8, salt []uint8, buff *[PASSLEN]uint8) {
	if len(buf) > 8 {
		buf = buf[:8]
	}

	// crypt.c: line 554
	x := uint8('A')
	if salt[0] != 0 {
		x = salt[0]
	}
	buff[0] = x

	// crypt.c: line 555
	Eswap0 := uint32(con_salt[x])

	// crypt.c: line 556
	x = uint8('A')
	if salt[1] != 0 {
		x = salt[1]
	}
	buff[1] = x

	// crypt: line 557
	Eswap1 := uint32(con_salt[x]) << 4

	// crypt: line 559
	key := desCBlock{}
	for idx, c := range buf {
		if c == 0 {
			break
		}

		key[idx] = c << 1
	}

	// crypt: line 568
	//log.Infof("fcrypt: key: %v", key)
	ks := desKeySchedule{}
	desSetKey(&key, &ks)
	/*
		for idx, each := range ks {
			log.Infof("fcrypt: (%d/%d) after desSetKey: ks: %v", idx, len(ks), each)
		}
	*/

	out := [2]uint32{}
	out[0], out[1] = body(&ks, Eswap0, Eswap1)

	//log.Infof("fcrypt: after body: out: %v", out)

	bb := [9]uint8{}
	b := &bb
	l2c(out[0], b[0:4])
	l2c(out[1], b[4:8])

	y := 0
	u := uint8(0x80)
	bb[8] = 0
	//log.Infof("fcrypt: bb: %v", bb)
	for i := 2; i < 13; i++ {
		c := 0
		for j := 0; j < 6; j++ {
			c <<= 1
			if bb[y]&u != 0 {
				c |= 1
			}
			u >>= 1
			if u == 0 {
				y++
				u = uint8(0x80)
			}
		}
		buff[i] = cov_2char[c]
		//log.Infof("fcrypt: (%d/%d) c: %v buff: %v", i, 13, c, buff[i])
	}
	buff[13] = 0
}

//cFcrypt
//
//This is the go-implementation of
//
//https://github.com/ptt/pttbbs/blob/master/common/sys/crypt.c
//commit: 6bdd36898bde207683a441cdffe2981e95de5b20
//
//Params
//
//	key: input-key
//	schedule: output-schedule
func desSetKey(key *desCBlock, schedule *desKeySchedule) {
	in := (*[8]uint8)(key)
	k := schedule

	c := c2l(in[:4])
	d := c2l(in[4:8])
	t := uint32(0)
	//log.Infof("desSetKey (1): to PermOp: d: %v c: %v", d, c)
	d, c = PermOp(d, c, 4, 0x0f0f0f0f)
	//log.Infof("desSetKey (1): after PermOp: d: %v c: %v", d, c)
	//log.Infof("desSetKey (2): to HPermOp: c: %v", c)
	c = HPermOp(c, -2, 0xcccc0000)
	//log.Infof("desSetKey (2): after HPermOp: c: %v", c)
	//log.Infof("desSetKey (3): to HPermOp: d: %v", d)
	d = HPermOp(d, -2, 0xcccc0000)
	//log.Infof("desSetKey (3): after HPermOp: d: %v", d)
	//log.Infof("desSetKey (4): to PermOp: d: %v c: %v", d, c)
	d, c = PermOp(d, c, 1, 0x55555555)
	//log.Infof("desSetKey (4): after PermOp: d: %v c: %v", d, c)
	//log.Infof("desSetKey (5): to PermOp: c: %v d: %v", c, d)
	c, d = PermOp(c, d, 8, 0x00ff00ff)
	//log.Infof("desSetKey (5): after PermOp: c: %v d: %v", c, d)
	//log.Infof("desSetKey (6): to PermOp: d: %v c: %v", d, c)
	d, c = PermOp(d, c, 1, 0x55555555)
	//log.Infof("desSetKey (6): after PermOp: d: %v c: %v", d, c)

	d = ((d & 0x000000ff) << 16) | (d & 0x0000ff00) |
		((d & 0x00ff0000) >> 16) | ((c & 0xf0000000) >> 4)

	c &= 0x0fffffff

	//log.Infof("desSetKey (7): to for-loop: c: %v d: %v", c, d)

	s := uint32(0)
	for i := 0; i < ITERATIONS; i++ {
		if shifts2[i] {
			{
				c = ((c >> 2) | (c << 26))
				d = ((d >> 2) | (d << 26))
			}
		} else {
			{
				c = ((c >> 1) | (c << 27))
				d = ((d >> 1) | (d << 27))
			}
		}

		//log.Infof("desSetKey (8) (%v/%v): to &= 0x0fffffff: c: %v d: %v", i, ITERATIONS, c, d)

		c &= 0x0fffffff
		d &= 0x0fffffff
		/* could be a few less shifts but I am to lazy at this
		 * point in time to investigate */
		s = skb[0][c&0x3f] |
			skb[1][((c>>6)&0x03)|((c>>7)&0x3c)] |
			skb[2][((c>>13)&0x0f)|((c>>14)&0x30)] |
			skb[3][((c>>20)&0x01)|((c>>21)&0x06)|
				((c>>22)&0x38)]
		t = skb[4][(d)&0x3f] |
			skb[5][((d>>7)&0x03)|((d>>8)&0x3c)] |
			skb[6][(d>>15)&0x3f] |
			skb[7][((d>>21)&0x0f)|((d>>22)&0x30)]
		//log.Infof("desSetKey (9) (%v/%v): to set k: s: %v t: %v", i, ITERATIONS, s, t)

		/* table contained 0213 4657 */
		k[2*i] = ((t << 16) | (s & 0x0000ffff)) & 0xffffffff
		s = ((s >> 16) | (t & 0xffff0000))

		s = (s << 4) | (s >> 28)
		k[2*i+1] = s & 0xffffffff

		//log.Infof("desSetKey (10) (%v/%v): after set k: (%v, %v)", i, ITERATIONS, k[2*i], k[2*i+1])
	}
}

func dEncrypt(L uint32, R uint32, S uint32, E0 uint32, E1 uint32, s []uint32, t uint32, u uint32) (retL uint32, retT uint32, retU uint32) {
	t = (R ^ (R >> 16))
	//log.Infof("dEncrypt (1): after t: R: %v t: %v", R, t)
	u = t & E0
	//log.Infof("dEncrypt (2): after u: t: %v E0: %v u: %v E1: %v", t, E0, u, E1)
	t = t & E1
	//log.Infof("dEncrypt (4): to u: t: %v u: %v R: %v S: %v s[S]: %v", t, u, R, S, s[S])
	u = (u ^ (u << 16)) ^ R ^ s[S]
	//log.Infof("dEncrypt (4): after u: u: %v", u)

	//log.Infof("dEncrypt (5): to t: t: %v R :%v s[S+1]: %v", t, R, s[S+1])
	t = (t ^ (t << 16)) ^ R ^ s[S+1]
	//log.Infof("dEncrypt (5): after t: t: %v", t)

	t = (t >> 4) | (t << 28)
	//log.Infof("dEncrypt (5): after t: t: %v", t)

	L ^= SPtrans[1][(t)&0x3f] |
		SPtrans[3][(t>>8)&0x3f] |
		SPtrans[5][(t>>16)&0x3f] |
		SPtrans[7][(t>>24)&0x3f] |
		SPtrans[0][(u)&0x3f] |
		SPtrans[2][(u>>8)&0x3f] |
		SPtrans[4][(u>>16)&0x3f] |
		SPtrans[6][(u>>24)&0x3f]
	//log.Infof("dEncrypt (6): after L: L: %v t: %v u: %v", L, t, u)

	return L, t, u
}

func body(ks *desKeySchedule, Eswap0 uint32, Eswap1 uint32) (uint32, uint32) {
	// crypt.c line: 618

	E0 := uint32(Eswap0)
	E1 := uint32(Eswap1)
	l := uint32(0)
	r := uint32(0)
	s := (*[32]uint32)(ks)
	t := uint32(0)
	u := uint32(0)

	//log.Infof("body: to for-loop: E0: %v E1: %v", E0, E1)
	for j := 0; j < 25; j++ {
		for i := uint32(0); i < ITERATIONS*2; i += 4 {
			//log.Infof("body (%v/%v)(%v/%v) (1): to dEncrypt: l: %v r: %v t: %v u: %v", j, 25, i, ITERATIONS*2, l, r, t, u)
			l, t, u = dEncrypt(l, r, i, E0, E1, s[:], t, u)
			//log.Infof("body (%v/%v)(%v/%v) (1): after dEncrypt: l: %v r: %v t: %v u: %v", j, 25, i, ITERATIONS*2, l, r, t, u)
			//log.Infof("body (%v/%v)(%v/%v) (2): to dEncrypt: l: %v r: %v t: %v u: %v", j, 25, i, ITERATIONS*2, l, r, t, u)
			r, t, u = dEncrypt(r, l, i+2, E0, E1, s[:], t, u)
			//log.Infof("body (%v/%v)(%v/%v) (2): after dEncrypt: l: %v r: %v t: %v u: %v", j, 25, i, ITERATIONS*2, l, r, t, u)
		}
		t = l
		l = r
		r = t

	}
	//log.Infof("body (3): after for-loop: l: %v r: %v t: %v u: %v", l, r, t, u)

	t = r
	r = (l >> 1) | (l << 31)
	l = (t >> 1) | (t << 31)

	l &= 0xffffffff
	r &= 0xffffffff

	//log.Infof("body (4): to PermOp: l: %v r: %v", l, r)
	r, l = PermOp(r, l, 1, 0x55555555)
	//log.Infof("body (4): after PermOp: l: %v r: %v", l, r)
	//log.Infof("body (5): to PermOp: l: %v r: %v", l, r)
	l, r = PermOp(l, r, 8, 0x00ff00ff)
	//log.Infof("body (5): after PermOp: l: %v r: %v", l, r)
	//log.Infof("body (6): to PermOp: l: %v r: %v", l, r)
	r, l = PermOp(r, l, 2, 0x33333333)
	//log.Infof("body (6): after PermOp: l: %v r: %v", l, r)
	//log.Infof("body (7): to PermOp: l: %v r: %v", l, r)
	l, r = PermOp(l, r, 16, 0x0000ffff)
	//log.Infof("body (7): after PermOp: l: %v r: %v", l, r)
	//log.Infof("body (8): to PermOp: l: %v r: %v", l, r)
	r, l = PermOp(r, l, 4, 0x0f0f0f0f)
	//log.Infof("body (8): after PermOp: l: %v r: %v", l, r)

	return l, r
}
