export const getJSTDate = () => {
  const now = new Date();
  const utc = now.getTime() + (now.getTimezoneOffset() * 60000);
  const jstOffset = 9 * 60 * 60000; // JSTはUTC+9
  const jstDate = new Date(utc + jstOffset);

  return jstDate.toISOString();
};

export const ToJSDate = (dateInput: Date | string | number): Date => {
  const date = new Date(dateInput);
  const utc = date.getTime() - (date.getTimezoneOffset() * 60000);
  const jstOffset = 9 * 60 * 60000; // JSTはUTC+9
  const jstDate = new Date(utc + jstOffset);

  return jstDate;
}
