export async function load({ params }) {
  const appName = params.appName;
  return { appName };
}
