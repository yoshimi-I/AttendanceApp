import { ActivityCalendar } from "@/components/githubCarender/githubCarender";
import CurrentTime from "../components/time/currentTime";
export default function Home() {
  const sampleData = [
    {
      day: "2023-01-01",
      activity: 5
    },
    {
      day: "2023-11-02",
      activity: 1
    }
  ]
  return (
    <main>
      <h1>Home</h1>
      <CurrentTime />
      <ActivityCalendar sampleData={sampleData}/>
    </main>
  );
}
