"use client";
import { useState, useEffect } from "react";
import { Box, Typography } from "@mui/material";

const dayOfWeek = ["日", "月", "火", "水", "木", "金", "土"];

const CurrentTime = () => {
  const [currentTime, setCurrentTime] = useState(new Date());
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    // コンポーネントがクライアントサイドでマウントされたことを確認
    setIsClient(true);
    setCurrentTime(new Date());

    const timer = setInterval(() => {
      setCurrentTime(new Date());
    }, 1000);

    // Cleanup timer when component is unmounted
    return () => clearInterval(timer);
  }, []);

  if (!isClient) {
    return null;
  }

  const year = currentTime.getFullYear();
  const month = currentTime.getMonth() + 1;
  const date = currentTime.getDate();
  const day = currentTime.getDay();

  const hours = currentTime.getHours();
  const minutes = currentTime.getMinutes();
  const seconds = currentTime.getSeconds();

  return (
    <Box style={{ fontFamily: "Roboto", textAlign: "center", padding: "40px" }}>
      <Typography variant="h4" style={{ fontWeight: 300 }}>
        {year}年{month}月{date}日 ({dayOfWeek[day]})
      </Typography>
      <Box display="flex" alignItems="baseline" justifyContent="center">
        <Typography variant="h1" style={{ fontWeight: 400, margin: "10px" }}>
          {hours < 10 ? "0" + hours : hours}
        </Typography>
        <span style={{ alignSelf: "center" }}>
          <Typography variant="h3" style={{ fontWeight: 600, margin: "10px" }}>
            :
          </Typography>
        </span>

        <Typography variant="h1" style={{ fontWeight: 400, margin: "10px" }}>
          {minutes < 10 ? "0" + minutes : minutes}
        </Typography>

        <Typography variant="h3" style={{ fontWeight: 300, margin: "20px" }}>
          {seconds < 10 ? "0" + seconds : seconds}
        </Typography>
      </Box>
    </Box>
  );
};

export default CurrentTime;
