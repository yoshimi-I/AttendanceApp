"use client";
import { useEffect, useState } from "react";
import CurrentTime from "../../components/time/currentTime";
import UserIcon from "../../components/userIcon/userIcon";
import { auth } from "../../lib/firebase/firebase";
import { useAuthState } from "react-firebase-hooks/auth";
import { useRouter } from "next/navigation";
import { ActivityCalendar } from "../../components/githubCarender/githubCarender";
import { AttendanceButtons } from "../../components/register/registerButtonList";

export default function Home() {
  const [user, loading, error] = useAuthState(auth);
  const router = useRouter();
  const [sampleData, setSampleData] = useState([]);
  const [userStatus, setUserStatus] = useState("");
  const year = new Date().getFullYear();
  const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

  useEffect(() => {
    if (loading) return;

    if (error) {
      console.error("Firebase auth error:", error);
      return;
    }

    if (!user) {
      router.push("/signin");
      return;
    }

    const fetchData = async () => {
      try {
        const response = await fetch(
          `${baseUrl}/study/allHistory/${user.uid}/${year}`
        );
        if (!response.ok) {
          throw new Error("Data fetching failed");
        }
        const data = await response.json();
        setSampleData(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    const fetchStatus = async () => {
      try {
        const response = await fetch(`${baseUrl}/user/status/${user.uid}`);
        if (!response.ok) {
          throw new Error("Data fetching failed");
        }
        const data = await response.json();
        setUserStatus(data.status);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
    fetchStatus();
  }, [user, loading, error, year, router]);

  if (loading) {
    return <div style={{ height: "100vh", backgroundColor: "white" }}></div>;
  }

  return (
    <main>
      <CurrentTime />
      <UserIcon />
      <AttendanceButtons userStatus={userStatus} />
      <ActivityCalendar sampleData={sampleData} />
    </main>
  );
}
