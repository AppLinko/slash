import { PlanType } from "@/types/proto/api/v1/subscription_service";

export const stringifyPlanType = (planType: PlanType = PlanType.FREE) => {
  if (planType === PlanType.FREE) {
    return "Enterprise";
  } else if (planType === PlanType.PRO) {
    return "Enterprise";
  } else if (planType === PlanType.ENTERPRISE) {
    return "Enterprise";
  } else {
    return "Unknown";
  }
};
