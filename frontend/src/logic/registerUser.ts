export const registerUser = async (firebaseUser) => {
  // APIエンドポイントのURL
  const apiUrl = "http://localhost:8080/user";

  console.log("registerUser function called with user:", firebaseUser);

  const userDTO = {
    name: firebaseUser.displayName, // Firebaseから取得したユーザー名
    email: firebaseUser.email, // Firebaseから取得したメールアドレス
    user_key: firebaseUser.uid, // Firebaseから取得したユーザーID
  };

  try {
    const response = await fetch(apiUrl, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userDTO),
    });
    
    console.log("Response received:", response);

    if (!response.ok) {
      throw new Error("Failed to register user");
    }

    const data = await response.json();
    console.log("User registered:", data);
  } catch (error) {
    console.error("Error registering user:", error);
  }
};
