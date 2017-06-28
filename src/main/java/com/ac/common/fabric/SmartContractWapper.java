package com.ac.common.fabric;

import com.ac.common.constant.SmartContractConstant;
import com.ac.common.fabric.model.ChainCodeResultModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.DefaultResourceLoader;
import org.springframework.core.io.ResourceLoader;
import org.springframework.stereotype.Service;

import java.io.File;

/**
 * Created by zhenchao.bi on 6/27/2017.
 */
@Service
public class SmartContractWapper {

    @Autowired
    private ChannelWapper channel;

    private ResourceLoader loader = new DefaultResourceLoader();


    public ChainCodeResultModel installHospitalSC() {

        try {
            //smartContract\hospital
            File scFile = loader.getResource("classpath:/smartContract/hospital").getFile();

            return channel.installChaincode(scFile, SmartContractConstant.Hospital.CHAINCODE_NAME, SmartContractConstant.Hospital.CHAINCODE_VERSION,
                    SmartContractConstant.Hospital.CHAINCODE_PATH, channel.getAllPeers());
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return null;
    }

    public ChainCodeResultModel installInsuranceSC() {

        try {
            //smartContract\hospital
            File scFile = loader.getResource("classpath:/smartContract/insurance").getFile();

            return channel.installChaincode(scFile, SmartContractConstant.Insurance.CHAINCODE_NAME, SmartContractConstant.Insurance.CHAINCODE_VERSION,
                    SmartContractConstant.Insurance.CHAINCODE_PATH, channel.getAllPeers());
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return null;
    }

    public ChainCodeResultModel installCustomerSC() {

        try {
            //smartContract\hospital
            File scFile = loader.getResource("classpath:/smartContract/customer").getFile();

            return channel.installChaincode(scFile, SmartContractConstant.Customer.CHAINCODE_NAME, SmartContractConstant.Customer.CHAINCODE_VERSION,
                    SmartContractConstant.Customer.CHAINCODE_PATH, channel.getAllPeers());
        } catch (Exception ex) {
            ex.printStackTrace();
        }

        return null;
    }

}
