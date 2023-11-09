"use client";
import { useEffect, useState } from 'react';
import CurrentTime from '../../components/time/currentTime';
import AttendanceButtons from '../../components/register/registerButtonList';
import UserIcon from '../../components/userIcon/userIcon';
import { auth } from "../firebase";
import { useAuthState } from "react-firebase-hooks/auth";
import { useRouter } from "next/navigation";


export default function Home() {
  const [user, loading, error] = useAuthState(auth);
  const router = useRouter();
  const [sampleData, setSampleData] = useState([]);
  const year = new Date().getFullYear();

  useEffect(() => {
    if (loading) return;

    if (error) {
      console.error("Firebase auth error:", error);
      return;
    }

    if (!user) {
      router.push('/signin');
      return;
    }

    const fetchData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/study/allHistory/${user.uid}/${year}`);
        if (!response.ok) {
          throw new Error('Data fetching failed');
        }
        const data = await response.json();
        setSampleData(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [user, loading, error, year, router]);

  if (loading) {
    // While loading, return a blank page or loading spinner
    return <div style={{ height: '100vh', backgroundColor: 'white' }}></div>;
  }

  return (
    <main>
      <CurrentTime />
      <UserIcon />
      <AttendanceButtons />
      <ActivityCalendar sampleData={sampleData} />
    </main>
  );
}


import { ActivityCalendar } from '../../components/githubCarender/githubCarender';
