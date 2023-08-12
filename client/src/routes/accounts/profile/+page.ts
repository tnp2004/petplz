import { PUBLIC_SERVER_URL } from '$env/static/public';
import { authenticated } from '../../../stores/auth';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ fetch }) => {
  const fetchAccount = async () => {
    try {
      const res = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
        method: 'GET',
        credentials: 'include',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json',
        }
      })

      const resOkay = res.ok;
      authenticated.set(resOkay);
      if (resOkay) {
        const data = await res.json()
        return data
      }
    } catch (error) {
      return error
    }
  }

  return {
    account: fetchAccount()
  }
}