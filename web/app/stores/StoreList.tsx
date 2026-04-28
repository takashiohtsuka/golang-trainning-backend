import type { StoreDetail } from "@/interfaces/store";
import Link from "next/link";

type Props = {
  stores: StoreDetail[];
};

export default function StoreList({ stores }: Props) {
  if (stores.length === 0) {
    return <p>店舗が見つかりませんでした。</p>;
  }

  return (
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>店舗名</th>
          <th>業種</th>
          <th>契約プラン</th>
          <th>公開状態</th>
          <th>有効</th>
          <th>地区</th>
          <th>都道府県</th>
          <th>地方</th>
        </tr>
      </thead>
      <tbody>
        {stores.map((store) => (
          <tr key={store.ID}>
            <td>{store.ID}</td>
            <td>
              <Link href={`/stores/${store.ID}`}>{store.Name}</Link>
            </td>
            <td>{store.BusinessTypeCode}</td>
            <td>{store.ContractPlanCode}</td>
            <td>{store.OpenStatus}</td>
            <td>{store.IsActive ? "有効" : "無効"}</td>
            <td>{store.DistrictName}</td>
            <td>{store.PrefectureName}</td>
            <td>{store.RegionName}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
