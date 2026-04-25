import type { StoreDetail as StoreDetailType } from "@/interfaces/store";

type Props = {
  store: StoreDetailType;
};

export default function StoreDetail({ store }: Props) {
  return (
    <>
      <h1>{store.Name}</h1>
      <ul>
        <li>ID: {store.ID}</li>
        <li>会社ID: {store.CompanyID}</li>
        <li>業種: {store.BusinessTypeCode}</li>
        <li>契約プラン: {store.ContractPlanCode}</li>
        <li>公開状態: {store.OpenStatus}</li>
        <li>有効: {store.IsActive ? "はい" : "いいえ"}</li>
        <li>地区: {store.DistrictName}</li>
        <li>都道府県: {store.PrefectureName}</li>
        <li>地方: {store.RegionName}</li>
      </ul>
    </>
  );
}
