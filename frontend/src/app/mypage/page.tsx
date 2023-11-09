"use client";
import { useEffect, useState } from 'react';
import { useForm } from 'react-hook-form';
import { ActivityCalendar } from '../../components/githubCarender/githubCarender';
import CurrentTime from '../../components/time/currentTime';
import AttendanceButtons from '../../components/register/registerButtonList';
import UserIcon from '../../components/userIcon/userIcon';
import LogoutDialog from '../../logic/logout';

export default function Home() {
  const { handleSubmit } = useForm();
  const [sampleData, setSampleData] = useState([]);
  const year = new Date().getFullYear();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/study/allHistory/hogehoge/${year}`);
        const data = await response.json();
        setSampleData(data);
        outPutData(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [year]);

  const outPutData = (data: any) => {
    console.log(data);
  };

  return (

    <main>
      <CurrentTime />
      <UserIcon/>
      <AttendanceButtons />
      <ActivityCalendar sampleData={sampleData} />

    </main>
  );
}
