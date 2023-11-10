import React from "react";
import Button from "@mui/material/Button";
import { getAuth } from "firebase/auth";

type ButtonType = "作業開始" | "作業終了" | "休憩開始" | "休憩終了";
const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

interface AttendanceButtonProps {
  type: ButtonType;
}

const AttendanceButton: React.FC<AttendanceButtonProps> = ({ type }) => {
  let color: "primary" | "secondary" | "inherit";
  let endpoint: string;

  switch (type) {
    case "作業開始":
      color = "primary";
      endpoint = "/study/activity/work/start";
      break;
    case "作業終了":
      color = "secondary";
      endpoint = "/study/activity/work/end";
      break;
    case "休憩開始":
      color = "inherit";
      endpoint = "/study/activity/break/start";
      break;
    case "休憩終了":
      color = "inherit";
      endpoint = "/study/activity/break/end";
      break;
    default:
      return null;
  }

  const handleClick = async () => {
    const auth = getAuth();
    const user = auth.currentUser;
    const userKey = user ? user.uid : null;

    if (!userKey) {
      console.error("User is not logged in");
      return;
    }

    const currentTime = new Date().toISOString();
    const url = `${baseUrl}${endpoint}`;

    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          user_key: userKey,
          time: currentTime,
        }),
      });

      if (!response.ok) {
        throw new Error("Request failed");
      }
    } catch (error) {
      alert(error.message);
    }
  };

  return (
    <Button
      variant="contained"
      color={color}
      size="large"
      sx={{
        borderRadius: "50%",
        width: "150px",
        height: "150px",
        margin: "2%",
      }}
      onClick={handleClick}
    >
      {type}
    </Button>
  );
};

export default AttendanceButton;
