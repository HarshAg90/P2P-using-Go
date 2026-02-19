# 📡 Decentralized Chat Application (React Native + Libp2p)

A fully decentralized peer-to-peer chat application built using **React Native** and **Libp2p**.  
This application enables direct device-to-device communication without relying on a centralized server.

---

## 🚀 Project Overview

This project demonstrates:

- ✅ Peer-to-peer networking using Libp2p  
- ✅ Local peer discovery using mDNS  
- ✅ Real-time message exchange  
- ⏳ Planned: End-to-End Encryption  
- ⏳ Planned: Decentralized storage (IPFS / OrbitDB)

The architecture eliminates centralized servers and allows devices on the same network to discover and communicate directly.

---

## 🏗️ Architecture

```
React Native (UI Layer)
        ↓
Libp2p Node (Networking Layer)
        ↓
mDNS (Peer Discovery)
        ↓
Direct P2P Communication
```

### Components

| Layer | Technology | Purpose |
|-------|------------|----------|
| Mobile App | React Native | Cross-platform UI |
| P2P Networking | Libp2p | Peer-to-peer communication |
| Peer Discovery | mDNS | Local network discovery |
| Encryption | (Planned) Signal Protocol | End-to-end encryption |
| Storage | (Planned) IPFS / OrbitDB | Decentralized persistence |

---

## 📦 Tech Stack

- React Native  
- Libp2p  
- @libp2p/websockets  
- @libp2p/tcp  
- @libp2p/mdns  
- @chainsafe/libp2p-noise  

---

## 🛠️ Installation

### 1️⃣ Clone the Repository

```bash
git clone https://github.com/yourusername/decentralized-chat-app.git
cd decentralized-chat-app
```

---

### 2️⃣ Install Dependencies

```bash
npm install
```

Install libp2p packages:

```bash
npm install libp2p @libp2p/websockets @libp2p/tcp @libp2p/mdns @chainsafe/libp2p-noise
```

---

### 3️⃣ iOS Setup (Only if building for iOS)

```bash
cd ios
pod install
cd ..
```

---

### 4️⃣ Run the Application

For Android:

```bash
npx react-native run-android
```

For iOS:

```bash
npx react-native run-ios
```

---

## 📱 How It Works

### 1️⃣ Node Initialization
Each device creates a Libp2p node on app launch.

### 2️⃣ Peer Discovery
mDNS allows automatic discovery of peers on the same local network.

### 3️⃣ Auto-Dial
When a peer is discovered, the node automatically connects.

### 4️⃣ Messaging
Messages are sent directly over a Libp2p stream between peers.

No backend server.  
No centralized database.  
Fully peer-driven communication.

---

## 🔐 Security (Planned Enhancement)

Upcoming improvements include:

- End-to-end encryption using Signal Protocol  
- Identity key generation per user  
- Secure session establishment  
- Message signing & verification  

---

## 📡 Decentralization Scope

Currently:

- Local network decentralized communication  
- Direct peer-to-peer message transfer  

Future improvements:

- Internet-wide peer discovery (Bootstrap nodes / DHT)  
- Offline message persistence via IPFS  
- Multi-peer broadcast  
- Group chat  
- Blockchain-based identity (optional)  

---

## 📂 Project Structure

```
/src
  ├── components
  ├── screens
  │     └── ChatScreen.js
  ├── services
  └── App.js
```

---

## 🧪 Testing Strategy

To test peer-to-peer functionality:

1. Connect two physical devices to the same WiFi network.  
2. Install and run the app on both devices.  
3. Devices should auto-discover each other.  
4. Send messages between devices.  

---

## ⚠️ Limitations (Current Version)

- Works only on same local network (mDNS scope)  
- No persistent message storage  
- No encryption layer yet  
- No offline queueing  

---

## 🎯 Roadmap

- [ ] End-to-end encryption (Signal Protocol)  
- [ ] IPFS integration for message storage  
- [ ] Global peer discovery (Kademlia DHT)  
- [ ] User identity system  
- [ ] Group chats  
- [ ] Push notifications  
- [ ] Media file sharing  
- [ ] Production-grade connection handling  

---

## 💡 Why Decentralized Chat?

Traditional messaging apps rely on centralized servers that:

- Store metadata  
- Control access  
- Can censor communication  
- Create single points of failure  

This project demonstrates how peer-to-peer networking eliminates those dependencies.

---

## 🤝 Contributing

Pull requests are welcome.  

For major changes, open an issue first to discuss improvements.

---

## 📜 License

MIT License

---

## 👨‍💻 Author

Harsh Agnihotri



--- 
> [Socket Artical](https://dzone.com/articles/socket-programming-in-go) | [Related Blog](https://medium.com/@jimsinjaradze/building-a-p2p-chat-application-in-go-a-learning-journey-8d7122897bf3#id_token=eyJhbGciOiJSUzI1NiIsImtpZCI6ImJiNDM0Njk1OTQ0NTE4MjAxNDhiMzM5YzU4OGFlZGUzMDUxMDM5MTkiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIyMTYyOTYwMzU4MzQtazFrNnFlMDYwczJ0cDJhMmphbTRsamRjbXMwMHN0dGcuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTUwNzk2MDI3MTA2NTM4MjM5MzEiLCJlbWFpbCI6ImhhcnNoLmFnbmlob3RyaTkwQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYmYiOjE3NDkwNjY1MDUsIm5hbWUiOiJIYXJzaCBBZ25paG90cmkiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jTHVQdHFLeTVBLUktMGlhd0M4bC0zbG9PY1g3cmJiejNONU5hM2YyR045UHdxcEk1UnZhQT1zOTYtYyIsImdpdmVuX25hbWUiOiJIYXJzaCIsImZhbWlseV9uYW1lIjoiQWduaWhvdHJpIiwiaWF0IjoxNzQ5MDY2ODA1LCJleHAiOjE3NDkwNzA0MDUsImp0aSI6IjAyNjRhMmE3YTgyNzM2OGVkMTk4ZWFhZTMzNTJhMmQ2YjA5MDlmNjAifQ.ZA7IdsqTTbyURYlUmVHWaGDRWS_wWPstXhmxFDY1fhoFwaCsJVu_ZbOASTq9VhmTCvxuNcgT_G3EjNEOirPNI6pgqDRhlYJqPz0z_3xgShhs-7Gh9GBUG16rrPbTUVTtAvm2JclfA5_fswXkOKBfq1dUDxkZkTzIXuERNQy9KcnyPSBMEw2j26pqrIttRxiUHZexZ5zDTKzOpkM2sl1ahkeKHlUqeH_xm79U5jREaXaX-2oJi83Rn3tMeh3a0oeDllS1YME4VX_HlTFWCeb07_OG_NWILYQmSsQaz15G1Rg0GnnXnkQj9_3GCLTMKgrLQYPusUosDt2773Fu7Qa6uA)

