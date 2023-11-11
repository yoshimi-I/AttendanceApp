"use client";
import { Box } from "@mui/system";
import AttendanceButton from "./registerButton";
import { useState } from "react";
import { Paper, Typography } from "@mui/material";

const AttendanceButtons = () => {
  const [userStatus, setUserStatus] = useState("");

  return (
    <>
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        style={{ marginBottom: "20px" }}
      >
        <Paper
          elevation={3} // 影の強さ
          style={{
            width: "400px",
            height: "120px",
            borderRadius: "100px", // 角の丸み
            color: "#862fa8",
            textAlign: "center",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <Typography
            variant="h3"
            style={{ fontWeight: "bold", fontFamily: 'Arial' }}
          >
            {userStatus}
          </Typography>
        </Paper>
      </Box>
      <Box
        display="flex"
        flexDirection="row"
        gap={2}
        alignItems="center"
        justifyContent="center"
        style={{ marginBottom: "40px" }}
      >
        <AttendanceButton type="作業開始" userStatus={setUserStatus} />
        <AttendanceButton type="作業終了" userStatus={setUserStatus} />
        <AttendanceButton type="休憩開始" userStatus={setUserStatus} />
        <AttendanceButton type="休憩終了" userStatus={setUserStatus} />
      </Box>
    </>
  );
};

export default AttendanceButtons;
