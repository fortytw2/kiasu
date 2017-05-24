export const NOTIFICATION_ADD = "NOTIFICATION_ADD";
export const NOTIFICATION_REMOVE = "NOTIFICATION_REMOVE";

export const NOTIFICATION_LEVEL_INFO = "NOTIFICATION_LEVEL_INFO";
export const NOTIFICATION_LEVEL_WARNING = "NOTIFICATION_LEVEL_WARNING";

export function addNotification(level, message) {
  return { type: NOTIFICATION_ADD, level: level, message: message };
}

export function removeNotification(message) {
  return { type: NOTIFICATION_REMOVE, message: message };
}
