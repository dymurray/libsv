package signature

// // SigningItem contains the metadata needed to sign a transaction.
// type SigningItem struct {
// 	PublicKeyHash string `json:"publicKeyHash"`
// 	SigHash       string `json:"sigHash"`
// 	PublicKey     string `json:"publicKey,omitempty"`
// 	Signature     string `json:"signature,omitempty"`
// }
//
// // SigningPayload type
// type SigningPayload []*SigningItem
//
// // NewSigningPayload creates a new SigningPayload.
// func NewSigningPayload() *SigningPayload {
// 	sp := make([]*SigningItem, 0)
// 	p := SigningPayload(sp)
// 	return &p
// }
//
// // AddItem appends a new SigningItem to the SigningPayload array.
// func (sp *SigningPayload) AddItem(publicKeyHash string, sigHash string) {
// 	si := &SigningItem{
// 		PublicKeyHash: publicKeyHash,
// 		SigHash:       sigHash,
// 	}
//
// 	*sp = append(*sp, si)
// }
//
// // GetSighashForInput function
// func GetSighashForInput(tx *transaction.Transaction, sighashType uint32, inputNumber uint32) string {
//
// 	in := tx.Inputs[inputNumber]
//
// 	getPrevoutHash := func(tx *transaction.Transaction) []byte {
// 		buf := make([]byte, 0)
//
// 		for _, in := range tx.Inputs {
// 			txid, _ := hex.DecodeString(in.PreviousTxID[:])
// 			buf = append(buf, utils.ReverseBytes(txid)...)
// 			oi := make([]byte, 4)
// 			binary.LittleEndian.PutUint32(oi, in.PreviousTxOutIndex)
// 			buf = append(buf, oi...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	getSequenceHash := func(tx *transaction.Transaction) []byte {
// 		buf := make([]byte, 0)
//
// 		for _, in := range tx.Inputs {
// 			oi := make([]byte, 4)
// 			binary.LittleEndian.PutUint32(oi, in.SequenceNumber)
// 			buf = append(buf, oi...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	getOutputsHash := func(tx *transaction.Transaction, n int32) []byte {
// 		buf := make([]byte, 0)
//
// 		if n == -1 {
// 			for _, out := range tx.Outputs {
// 				buf = append(buf, out.GetBytesForSigHash()...)
// 			}
// 		} else {
// 			buf = append(buf, tx.Outputs[n].GetBytesForSigHash()...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	hashPrevouts := make([]byte, 32)
// 	hashSequence := make([]byte, 32)
// 	hashOutputs := make([]byte, 32)
//
// 	if sighashType&transaction.SighashAnyoneCanPay == 0 {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashPrevouts = getPrevoutHash(transaction)
// 	}
//
// 	if sighashType&transaction.SighashAnyoneCanPay == 0 &&
// 		(sighashType&31) != transaction.SighashSingle &&
// 		(sighashType&31) != transaction.SighashNone {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashSequence = getSequenceHash(transaction)
// 	}
//
// 	if (sighashType&31) != transaction.SighashSingle && (sighashType&31) != transaction.SighashNone {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashOutputs = getOutputsHash(transaction, -1)
// 	} else if (sighashType&31) == transaction.SighashSingle && inputNumber < uint32(len(transaction.Outputs)) {
// 		// This will *not* be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashOutputs = getOutputsHash(transaction, int32(inputNumber))
// 	}
//
// 	buf := make([]byte, 0)
//
// 	// Version
// 	v := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(v, transaction.Version)
// 	buf = append(buf, v...)
//
// 	// Input prevouts/nSequence (none/all, depending on flags)
// 	buf = append(buf, hashPrevouts...)
// 	buf = append(buf, hashSequence...)
//
// 	//  outpoint (32-byte hash + 4-byte little endian)
// 	txid, _ := hex.DecodeString(in.PreviousTxID[:])
// 	buf = append(buf, utils.ReverseBytes(txid)...)
// 	oi := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(oi, in.PreviousTxOutIndex)
// 	buf = append(buf, oi...)
//
// 	// scriptCode of the input (serialized as scripts inside CTxOuts)
// 	buf = append(buf, utils.VarInt(uint64(len(*in.PreviousTxScript)))...)
// 	buf = append(buf, *in.PreviousTxScript...)
//
// 	// value of the output spent by this input (8-byte little endian)
// 	sat := make([]byte, 8)
// 	binary.LittleEndian.PutUint64(sat, in.PreviousTxSatoshis)
// 	buf = append(buf, sat...)
//
// 	// nSequence of the input (4-byte little endian)
// 	seq := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(seq, in.SequenceNumber)
// 	buf = append(buf, seq...)
//
// 	// Outputs (none/one/all, depending on flags)
// 	buf = append(buf, hashOutputs...)
//
// 	// Locktime
// 	lt := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(lt, transaction.Locktime)
// 	buf = append(buf, lt...)
//
// 	// sighashType
// 	// writer.writeUInt32LE(sighashType >>> 0)
// 	st := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(st, sighashType>>0)
// 	buf = append(buf, st...)
// 	ret := crypto.Sha256d(buf)
// 	return hex.EncodeToString(utils.ReverseBytes(ret))
// }
//
// // GetSighashForInputValidation comment todo
// func GetSighashForInputValidation(transaction *transaction.Transaction, sighashType uint32, inputNumber uint32, previousTxOutIndex uint32, previousTxSatoshis uint64, previousTxScript *script.Script) string {
//
// 	in := transaction.Inputs[inputNumber]
//
// 	getPrevoutHash := func(tx *transaction.Transaction) []byte {
// 		buf := make([]byte, 0)
//
// 		for _, in := range tx.Inputs {
// 			txid, _ := hex.DecodeString(in.PreviousTxID[:])
// 			buf = append(buf, utils.ReverseBytes(txid)...)
// 			oi := make([]byte, 4)
// 			binary.LittleEndian.PutUint32(oi, previousTxOutIndex)
// 			buf = append(buf, oi...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	getSequenceHash := func(tx *transaction.Transaction) []byte {
// 		buf := make([]byte, 0)
//
// 		for _, in := range tx.Inputs {
// 			oi := make([]byte, 4)
// 			binary.LittleEndian.PutUint32(oi, in.SequenceNumber)
// 			buf = append(buf, oi...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	getOutputsHash := func(tx *transaction.Transaction, n int32) []byte {
// 		buf := make([]byte, 0)
//
// 		if n == -1 {
// 			for _, out := range tx.Outputs {
// 				buf = append(buf, out.GetBytesForSigHash()...)
// 			}
// 		} else {
// 			buf = append(buf, tx.Outputs[n].GetBytesForSigHash()...)
// 		}
//
// 		return crypto.Sha256d(buf)
// 	}
//
// 	hashPrevouts := make([]byte, 32)
// 	hashSequence := make([]byte, 32)
// 	hashOutputs := make([]byte, 32)
//
// 	if sighashType&transaction.SighashAnyoneCanPay == 0 {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashPrevouts = getPrevoutHash(transaction)
// 	}
//
// 	if sighashType&transaction.SighashAnyoneCanPay == 0 &&
// 		(sighashType&31) != transaction.SighashSingle &&
// 		(sighashType&31) != transaction.SighashNone {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashSequence = getSequenceHash(transaction)
// 	}
//
// 	if (sighashType&31) != transaction.SighashSingle && (sighashType&31) != transaction.SighashNone {
// 		// This will be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashOutputs = getOutputsHash(transaction, -1)
// 	} else if (sighashType&31) == transaction.SighashSingle && inputNumber < uint32(len(transaction.Outputs)) {
// 		// This will *not* be executed in the usual BSV case (where sighashType = SighashAllForkID)
// 		hashOutputs = getOutputsHash(transaction, int32(inputNumber))
// 	}
//
// 	buf := make([]byte, 0)
//
// 	// Version
// 	v := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(v, transaction.Version)
// 	buf = append(buf, v...)
//
// 	// Input prevouts/nSequence (none/all, depending on flags)
// 	buf = append(buf, hashPrevouts...)
// 	buf = append(buf, hashSequence...)
//
// 	//  outpoint (32-byte hash + 4-byte little endian)
// 	txid, _ := hex.DecodeString(in.PreviousTxID[:])
// 	buf = append(buf, utils.ReverseBytes(txid)...)
// 	oi := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(oi, previousTxOutIndex)
// 	buf = append(buf, oi...)
//
// 	// scriptCode of the input (serialized as scripts inside CTxOuts)
// 	buf = append(buf, utils.VarInt(uint64(len(*previousTxScript)))...)
// 	buf = append(buf, *previousTxScript...)
//
// 	// value of the output spent by this input (8-byte little endian)
// 	sat := make([]byte, 8)
// 	binary.LittleEndian.PutUint64(sat, previousTxSatoshis)
// 	buf = append(buf, sat...)
//
// 	// nSequence of the input (4-byte little endian)
// 	seq := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(seq, in.SequenceNumber)
// 	buf = append(buf, seq...)
//
// 	// Outputs (none/one/all, depending on flags)
// 	buf = append(buf, hashOutputs...)
//
// 	// Locktime
// 	lt := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(lt, transaction.Locktime)
// 	buf = append(buf, lt...)
//
// 	// sighashType
// 	// writer.writeUInt32LE(sighashType >>> 0)
// 	st := make([]byte, 4)
// 	binary.LittleEndian.PutUint32(st, sighashType>>0)
// 	buf = append(buf, st...)
// 	ret := crypto.Sha256d(buf)
// 	return hex.EncodeToString(utils.ReverseBytes(ret))
// }