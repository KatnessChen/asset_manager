-- +migrate Up
-- Rename 'amount' column to 'value'
ALTER TABLE asset_records
CHANGE COLUMN amount unit_price DECIMAL(15, 2) NOT NULL;

-- Add 'unit' column
ALTER TABLE asset_records
ADD COLUMN unit DECIMAL(15, 2) NOT NULL;

-- Add 'unit_cost' column
ALTER TABLE asset_records
ADD COLUMN unit_cost DECIMAL(15, 2) NOT NULL;