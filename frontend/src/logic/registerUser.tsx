export const registerUser = async (firebaseUser) => {
  // APIエンドポイントのURL
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  const userDTO = {
    name: firebaseUser.displayName, // Firebaseから取得したユーザー名
    email: firebaseUser.email, // Firebaseから取得したメールアドレス
    user_key: firebaseUser.uid, // Firebaseから取得したユーザーID
  };

  try {
    const response = await fetch(`${baseUrl}/user`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userDTO),
    });

    if (!response.ok) {
      throw new Error("Failed to register user");
    }
  } catch (error) {
    console.error("Error registering user:", error);
  }
};
