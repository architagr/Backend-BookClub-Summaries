# Notes for Chapter 3: Service Discovery

[Link to the recorded session on YouTube](https://www.youtube.com/live/wasYCK2oYd8?si=EEPbwoyT0K8Y7pD-)

## Service Discovery — The Missing Puzzle Piece of Microservices

Imagine trying to invite friends to a party without knowing where they live. Hardcoding their addresses might work once, but if they move? You're partying alone. That’s _exactly_ the problem we face when services talk to each other without service discovery.

In a microservices setup, hardcoding service addresses is a recipe for disaster. If a service instance moves, crashes, or scales, every dependent service must update its configuration. Manually. Painful, right?

### ❌ Hardcoded Service Addresses: The Two Big Problems

1. **Scalability Issues** – You can’t easily add or remove service instances.
2. **Reliability Problems** – If a hardcoded service crashes or becomes unreachable, the dependent services are left hanging.

These are classic **service discovery problems**.

## ✅ What Does Service Discovery Solve?

1. Automatically finds available service instances.
2. Adds/removes service addresses without manual intervention.
3. Detects and avoids unhealthy or crashed instances.

To solve this, we need:

- A **Registry** (a central database of all live service instances)
- A **Discovery Model** (how clients interact with the registry)
- **Health Monitoring** (to know who’s alive and who’s just pretending)

## 📦 The Service Registry

Think of this as a "live contacts list" for all your services.

A service registry’s responsibilities:

- Register new service instances
- Deregister them when they go offline
- Return a list of live instances for a given service

Examples: **_Consul, etcd, Zookeeper, Kubernetes_**

## 🔀 Discovery Models

1. **Client-Side Service Discovery**

   Here, each client talks directly to the registry to fetch the list of service instances and chooses one (maybe using round-robin, random, etc.).

   **🧠 Analogy:** It’s like checking Google Maps for nearby pizza places and choosing one yourself.

   **📦 Example:** In our inventory app, the gateway queries the registry for catalogue-service instances and load-balances requests between them.

2. **Server-Side Service Discovery**

   In this model, clients just call a load balancer, which internally talks to the registry and forwards the request to an appropriate service instance.

   **🧠 Analogy:** You call customer support, and an IVR routes you to an available agent.

   👍 This is more convenient for clients, as they don’t need discovery or load-balancing logic. It’s abstracted away.

## 🩺 Service Health Monitoring

We don’t just need to know where services are—we also need to know if they’re alive.

There are two common models:

1. **Pull Model**

   The registry checks the health of each instance at regular intervals.

   **🧠 Analogy:** Your mom texting every night to ask, “Did you eat?”

2. **Push Model**

   Each service sends a heartbeat to the registry at regular intervals to say, “I’m still breathing.”

   **🧠 Analogy:** A criminal checking in at the police station to prove they haven’t skipped town.

## 🛠️ Tools That Help

All of these support service discovery (and are written in Go! ❤️):

- **HashiCorp Consul** – Full-blown service discovery + health checks
- **Kubernetes** – Built-in discovery with services/endpoints
- **etcd** – A consistent key-value store used by Kubernetes
- **Apache Zookeeper** – Battle-tested in large-scale distributed systems
