package server

import "github.com/irmine/goraklib/protocol"

func (manager *SessionManager) HandleUnconnectedMessage(packetInterface protocol.IPacket, session *Session) {
	if session.IsOpened() {
		return
	}
	switch packet := packetInterface.(type) {
	case *protocol.UnconnectedPing:
		var pong = protocol.NewUnconnectedPong()

		pong.PingTime = manager.server.GetRunTime()
		pong.ServerId = manager.server.GetServerId()
		pong.ServerName = manager.server.GetServerName()
		pong.ServerProtocol = manager.server.GetMinecraftProtocol()
		pong.ServerVersion = manager.server.GetMinecraftVersion()
		pong.Motd = manager.server.GetMotd()
		pong.DefaultGameMode = manager.server.GetDefaultGameMode()
		pong.MaximumPlayers = manager.server.GetMaxConnectedSessions()
		pong.OnlinePlayers = manager.server.GetConnectedSessionCount()

		session.SendUnconnectedPacket(pong)

	case *protocol.OpenConnectionRequest1:
		var response = protocol.NewOpenConnectionResponse1()

		response.ServerId = manager.server.GetServerId()
		response.MtuSize = packet.MtuSize

		response.Security = manager.server.IsSecure()

		session.SendUnconnectedPacket(response)

	case *protocol.OpenConnectionRequest2:
		var response = protocol.NewOpenConnectionResponse2()

		response.ClientPort = uint16(session.port)
		response.ClientAddress = session.address
		response.MtuSize = packet.MtuSize
		response.ServerId = manager.server.GetServerId()

		session.mtuSize = response.MtuSize

		response.UseEncryption = manager.server.useEncryption

		session.SendUnconnectedPacket(response)

		session.Open()
	}
}
