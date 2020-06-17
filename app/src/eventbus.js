import Vue from 'vue';
export const EventBus = new Vue();

export const AddNewCardEvent  = "addNewCard";
export const UpdatedCardEvent = "UpdatedCardEvent";
export const DeleteCardEvent  = "DeleteCardEvent";