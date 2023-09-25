db.createUser(
  {
    user: "user123",
    pwd: "pass123",
    roles: [
      {
        role: "readWrite",
        db: "synapsis_challenge"
      }
    ]
  }
);