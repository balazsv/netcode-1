package netcode

import "errors"

var (
	ErrChallengePacketDecryptedDataSize  = errors.New("ignored connection challenge packet. decrypted packet data is wrong size")
	ErrChallengePacketTokenData          = errors.New("error reading challenge token data")
	ErrChallengePacketTokenSequence      = errors.New("error reading challenge token sequence")
	ErrClientNotConnected                = errors.New("client not connected, unable to send packet")
	ErrDecryptPacketPayloadTooSmall      = errors.New("ignored encrypted packet. encrypted payload is too small")
	ErrDeniedPacketDecryptedDataSize     = errors.New("ignored connection denied packet. decrypted packet data is wrong size")
	ErrDisconnectPacketDecryptedDataSize = errors.New("ignored connection denied packet. decrypted packet data is wrong size")
	ErrEmptyPacketKey                    = errors.New("empty packet key")
	ErrEmptyServers                      = errors.New("empty servers")
	ErrEncryptedPacketBufferTooSmall     = errors.New("ignored encrypted packet. buffer is too small for sequence bytes + encryption mac")
	ErrEncryptedPacketSequenceOutOfRange = errors.New("ignored encrypted packet. sequence bytes is out of range [1,8]")
	ErrEncryptedPacketTooSmall           = errors.New("ignored encrypted packet. packet is too small to be valid")
	ErrExceededServerNumber              = errors.New("invalid server address, exceeded # of servers")
	ErrExpiredTokenTimestamp             = errors.New("expire timestamp is > create timestamp")
	ErrInvalidBufferLength               = errors.New("invalid buffer length")
	ErrInvalidBufferSize                 = errors.New("invalid buffer size written")
	ErrInvalidIPAddress                  = errors.New("invalid ip address")
	ErrInvalidPacket                     = errors.New("invalid packet")
	ErrInvalidPort                       = errors.New("invalid port")
	ErrInvalidSequenceBytes              = errors.New("invalid sequence bytes, must be between [1-8]")
	ErrInvalidTokenPrivateByteSize       = errors.New("error in encrypt invalid token private byte size")
	ErrKeepAlivePacketClientIndex        = errors.New("error reading keepalive client index")
	ErrKeepAlivePacketDecryptedDataSize  = errors.New("ignored connection keep alive packet. decrypted packet data is wrong size")
	ErrKeepAlivePacketMaxClients         = errors.New("error reading keepalive max clients")
	ErrPacketHandlerBeforeListen         = errors.New("packet handler must be set before calling listen")
	ErrPacketSizeMax                     = errors.New("packet size was > maxBytes")
	ErrPayloadPacketTooLarge             = errors.New("payload packet: payload is too large")
	ErrPayloadPacketTooSmall             = errors.New("payload packet: payload is too small")
	ErrReadingUserData                   = errors.New("error reading user data")
	ErrRequestBadPacketLength            = errors.New("request packet: bad packet length")
	ErrRequestPacketBadProtocolId        = errors.New("request packet: wrong protocol id")
	ErrRequestPacketBadVersionInfo       = errors.New("request packet: bad version info did not match")
	ErrRequestPacketBadVersionInfoBytes  = errors.New("request packet: bad version info invalid bytes returned")
	ErrRequestPacketBufferInvalidLength  = errors.New("invalid length of packet buffer read")
	ErrRequestPacketConnectTokenExpired  = errors.New("request packet: connect token expired")
	ErrRequestPacketNoPrivateKey         = errors.New("request packet: no private key")
	ErrRequestPacketTypeNotAllowed       = errors.New("request packet: packet type is not allowed")
	ErrResponsePacketDecryptedDataSize   = errors.New("challenge response packet: decrypted packet data is wrong size")
	ErrResponsePacketTokenData           = errors.New("error reading challenge token data")
	ErrResponsePacketTokenSequence       = errors.New("error reading challenge token sequence")
	ErrServerNotRunning                  = errors.New("server is not running")
	ErrServerShutdown                    = errors.New("server shutdown")
	ErrSocketZeroRecv                    = errors.New("socket error: 0 byte length recv'd")
	ErrTooManyServers                    = errors.New("too many servers")
	ErrUnknownIPAddress                  = errors.New("unknown ip address")
	ErrWriteClosedSocket                 = errors.New("unable to write, socket has been closed")
)
