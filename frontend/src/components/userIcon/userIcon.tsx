import { useEffect, useState } from "react";
import { User, onAuthStateChanged } from "firebase/auth";
import Image from "next/image";
import { auth } from "../../lib/firebase/firebase";
import { AccountCircle } from "@mui/icons-material";
import LogoutDialog from "../logout/logoutButton";

const UserIcon = () => {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    // Firebaseの認証状態の変更を監視
    const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
      setUser(currentUser);
    });

    return () => unsubscribe();
  }, []);

  if (!user) {
    return null;
  }

  return (
    <div
      style={{
        position: "absolute",
        top: 0,
        right: 0,
        padding: "10px",
        display: "flex",
        alignItems: "center",
      }}
    >
      {user.photoURL ? (
        <Image
          src={user.photoURL}
          alt="User Icon"
          width={40}
          height={40}
          style={{
            borderRadius: "50%",
            cursor: "pointer",
            marginRight: "10px",
          }}
        />
      ) : (
        <AccountCircle
          style={{ fontSize: 40, cursor: "pointer", marginRight: "10px" }}
        />
      )}
      <LogoutDialog />
    </div>
  );
};

export default UserIcon;
