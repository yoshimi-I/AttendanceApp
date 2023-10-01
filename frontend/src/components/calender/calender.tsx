import dynamic from 'next/dynamic';

// ActivityCalendarをクライアントサイドでのみ動的にインポート
const DynamicActivityCalendar = dynamic(
  () => import('activity-calendar-react').then(mod => mod.ActivityCalendar),
  {
    ssr: false, // SSRを無効にする
    loading: () => <p>Loading...</p> // オプション: ロード中の表示内容
  }
);

function GithubCalendar() {
  return (
    <div>
      <DynamicActivityCalendar />
    </div>
  );
}

export default GithubCalendar;
