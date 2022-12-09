export default function IsDiffObj(obj: any, origin: any) {
  if (obj && origin) {
    return  Object.keys(obj).some(x=>obj[x] !== origin[x])
  }

  return true;
}
