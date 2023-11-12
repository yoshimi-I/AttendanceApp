export const getJSTDate = () => {
  const now = new Date();
  const jstDate = now.toLocaleString("en-US", { timeZone: "Asia/Tokyo" });
  return new Date(jstDate);
};

export const ToJSDate = (dateInput: Date | string | number): Date => {
  const date = new Date(dateInput);
  const utc = date.getTime() - date.getTimezoneOffset() * 60000;
  const jstOffset = 9 * 60 * 60000; // JSTはUTC+9
  const jstDate = new Date(utc + jstOffset);
  return jstDate;
};

export const StrToJSDate = (dateInput: string): Date => {
  const time = new Date();
  const [hours, minutes] = dateInput.split(":").map(Number);
  time.setHours(hours, minutes, 0, 0); // 秒とミリ秒は0に設定
  return time;
}
